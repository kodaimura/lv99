package executor

type CodeExecRequest struct {
	CodeDef  string `json:"code_def"`
	CodeCall string `json:"code_call"`
}

type CodeExecResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}
