import json
import os
import re
import resource
import subprocess
import sys
import tempfile
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from pathlib import Path

HOST = "0.0.0.0"
PORT = int(os.getenv("CODE_RUNNER_PORT", "8080"))
TIMEOUT_SECONDS = float(os.getenv("CODE_RUNNER_TIMEOUT_SECONDS", "5"))
MAX_BODY_BYTES = int(os.getenv("CODE_RUNNER_MAX_BODY_BYTES", "131072"))
MAX_SOURCE_BYTES = int(os.getenv("CODE_RUNNER_MAX_SOURCE_BYTES", "65536"))
MAX_OUTPUT_BYTES = int(os.getenv("CODE_RUNNER_MAX_OUTPUT_BYTES", "65536"))
MEMORY_LIMIT_BYTES = int(os.getenv("CODE_RUNNER_MEMORY_LIMIT_BYTES", str(256 * 1024 * 1024)))
CPU_LIMIT_SECONDS = int(os.getenv("CODE_RUNNER_CPU_LIMIT_SECONDS", "3"))
FILE_SIZE_LIMIT_BYTES = int(os.getenv("CODE_RUNNER_FILE_SIZE_LIMIT_BYTES", str(1024 * 1024)))

TRACEBACK_FILE_RE = re.compile(r'  File "[^"]+", line (\d+), in ([^\n]+)')
RUNNER_SOURCE = r'''
import ast
import linecache
import sys
import traceback


class LastExpressionPrinter(ast.NodeTransformer):
    def visit_Module(self, node):
        self.generic_visit(node)
        if not node.body or not isinstance(node.body[-1], ast.Expr):
            return node

        expr = node.body[-1]
        result_name = "__lv99_last_result"
        assign = ast.Assign(
            targets=[ast.Name(id=result_name, ctx=ast.Store())],
            value=expr.value,
        )
        check = ast.If(
            test=ast.Compare(
                left=ast.Name(id=result_name, ctx=ast.Load()),
                ops=[ast.IsNot()],
                comparators=[ast.Constant(value=None)],
            ),
            body=[
                ast.Expr(
                    value=ast.Call(
                        func=ast.Name(id="print", ctx=ast.Load()),
                        args=[ast.Name(id=result_name, ctx=ast.Load())],
                        keywords=[],
                    )
                )
            ],
            orelse=[],
        )
        ast.copy_location(assign, expr)
        ast.copy_location(check, expr)
        ast.copy_location(check.body[0], expr)
        node.body[-1:] = [assign, check]
        return ast.fix_missing_locations(node)


def format_user_exception(exc, code_path):
    source = code_path.read_text(encoding="utf-8")
    linecache.cache[str(code_path)] = (len(source), None, source.splitlines(True), str(code_path))

    print("Traceback (most recent call last):", file=sys.stderr)
    user_frames = []
    tb = exc.__traceback__
    while tb is not None:
        frame = tb.tb_frame
        if frame.f_code.co_filename == str(code_path):
            user_frames.append((tb.tb_lineno, frame.f_code.co_name))
        tb = tb.tb_next

    if isinstance(exc, SyntaxError) and exc.lineno is not None:
        user_frames.append((exc.lineno, "<module>"))

    for lineno, name in user_frames:
        print(f'  File "{code_path}", line {lineno}, in {name}', file=sys.stderr)
        line = linecache.getline(str(code_path), lineno).rstrip()
        if line:
            print(f"    {line}", file=sys.stderr)

    print("".join(traceback.format_exception_only(type(exc), exc)), end="", file=sys.stderr)


def main():
    from pathlib import Path

    code_path = Path(sys.argv[1])
    source = code_path.read_text(encoding="utf-8")
    try:
        tree = ast.parse(source, filename=str(code_path), mode="exec")
        tree = LastExpressionPrinter().visit(tree)
        code = compile(tree, str(code_path), "exec")
        exec(code, {"__name__": "__main__"})
    except SystemExit:
        raise
    except BaseException as exc:
        format_user_exception(exc, code_path)
        sys.exit(1)


if __name__ == "__main__":
    main()
'''


def truncate_text(value: str, limit: int = MAX_OUTPUT_BYTES) -> str:
    encoded = value.encode("utf-8", errors="replace")
    if len(encoded) <= limit:
        return value
    clipped = encoded[:limit].decode("utf-8", errors="ignore")
    return clipped + "\n... output truncated ..."


def normalize_output(stdout: str) -> str:
    return stdout.rstrip("\r\n")


def clean_traceback(stderr: str) -> str:
    # Keep the existing UI-friendly format: hide the temporary filename but keep line numbers.
    return TRACEBACK_FILE_RE.sub(r"  line \1, in \2", stderr)


def limit_child_process() -> None:
    safe_setrlimit(resource.RLIMIT_CPU, (CPU_LIMIT_SECONDS, CPU_LIMIT_SECONDS + 1))
    safe_setrlimit(resource.RLIMIT_AS, (MEMORY_LIMIT_BYTES, MEMORY_LIMIT_BYTES))
    safe_setrlimit(resource.RLIMIT_FSIZE, (FILE_SIZE_LIMIT_BYTES, FILE_SIZE_LIMIT_BYTES))
    if hasattr(resource, "RLIMIT_NPROC"):
        safe_setrlimit(resource.RLIMIT_NPROC, (64, 64))


def safe_setrlimit(resource_name: int, limits: tuple[int, int]) -> None:
    try:
        resource.setrlimit(resource_name, limits)
    except (OSError, ValueError):
        pass


def build_source(code_def: str, code_call: str) -> str:
    code_def = code_def.rstrip()
    code_call = code_call.lstrip()
    if code_def == "":
        return code_call
    if code_call == "":
        return code_def
    return code_def + "\n" + code_call


def execute_python(code_def: str, code_call: str) -> dict[str, str]:
    source = build_source(code_def, code_call)
    if len(source.encode("utf-8")) > MAX_SOURCE_BYTES:
        return {"output": "", "error": "ValueError: submitted code is too large"}

    with tempfile.TemporaryDirectory(prefix="lv99-code-") as tmp_dir:
        code_path = Path(tmp_dir) / "main.py"
        runner_path = Path(tmp_dir) / "runner.py"
        code_path.write_text(source, encoding="utf-8")
        runner_path.write_text(RUNNER_SOURCE, encoding="utf-8")

        env = {
            "HOME": tmp_dir,
            "LANG": "C.UTF-8",
            "LC_ALL": "C.UTF-8",
            "PATH": "/usr/local/bin:/usr/bin:/bin",
            "PYTHONIOENCODING": "utf-8",
            "PYTHONDONTWRITEBYTECODE": "1",
        }

        try:
            completed = subprocess.run(
                [sys.executable, "-I", str(runner_path), str(code_path)],
                cwd=tmp_dir,
                env=env,
                stdin=subprocess.DEVNULL,
                capture_output=True,
                text=True,
                timeout=TIMEOUT_SECONDS,
                preexec_fn=limit_child_process,
            )
        except subprocess.TimeoutExpired as exc:
            output = exc.stdout if isinstance(exc.stdout, str) else ""
            return {
                "output": truncate_text(output),
                "error": f"TimeoutError: code execution exceeded {TIMEOUT_SECONDS:g} seconds",
            }

    return {
        "output": truncate_text(normalize_output(completed.stdout)),
        "error": truncate_text(clean_traceback(completed.stderr)),
    }


class Handler(BaseHTTPRequestHandler):
    server_version = "lv99-code-runner/1.0"

    def do_GET(self) -> None:
        if self.path == "/health":
            self.write_json(200, {"status": "ok"})
            return
        self.write_json(404, {"error": "not found"})

    def do_POST(self) -> None:
        if self.path != "/exec/python":
            self.write_json(404, {"error": "not found"})
            return

        try:
            content_length = int(self.headers.get("Content-Length", "0"))
        except ValueError:
            self.write_json(400, {"error": "invalid content length"})
            return

        if content_length <= 0 or content_length > MAX_BODY_BYTES:
            self.write_json(413, {"error": "request body is too large"})
            return

        try:
            payload = json.loads(self.rfile.read(content_length).decode("utf-8"))
        except (UnicodeDecodeError, json.JSONDecodeError):
            self.write_json(400, {"error": "invalid json"})
            return

        code_def = payload.get("code_def", "")
        code_call = payload.get("code_call", "")
        if not isinstance(code_def, str) or not isinstance(code_call, str):
            self.write_json(400, {"error": "code_def and code_call must be strings"})
            return

        self.write_json(200, execute_python(code_def, code_call))

    def log_message(self, fmt: str, *args: object) -> None:
        print(f"{self.address_string()} - {fmt % args}", file=sys.stderr, flush=True)

    def write_json(self, status: int, body: dict[str, str]) -> None:
        encoded = json.dumps(body, ensure_ascii=False).encode("utf-8")
        self.send_response(status)
        self.send_header("Content-Type", "application/json; charset=utf-8")
        self.send_header("Content-Length", str(len(encoded)))
        self.end_headers()
        self.wfile.write(encoded)


if __name__ == "__main__":
    server = ThreadingHTTPServer((HOST, PORT), Handler)
    print(f"code runner listening on {HOST}:{PORT}", flush=True)
    server.serve_forever()
