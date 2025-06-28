package executor

type CodeExecutor interface {
	Execute(CodeExecRequest) (CodeExecResponse, error)
}
