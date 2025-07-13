package core

type CodeExecutorI interface {
	Execute(CodeExecRequest) (CodeExecResponse, error)
}

type CodeExecRequest struct {
	CodeDef  string `json:"code_def"`
	CodeCall string `json:"code_call"`
}

type CodeExecResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

var CodeExecutor CodeExecutorI = &noopCodeExecutor{}

func SetCodeExecutor(ce CodeExecutorI) {
	CodeExecutor = ce
}

type noopCodeExecutor struct{}

func (ce *noopCodeExecutor) Execute(CodeExecRequest) (CodeExecResponse, error) {
	return CodeExecResponse{}, nil
}
