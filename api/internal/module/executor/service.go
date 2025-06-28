package executor

type Service interface {
	Execute(in CodeExecRequest) (CodeExecResponse, error)
}

type service struct {
	executor CodeExecutor
}

func NewService(executor CodeExecutor) Service {
	return &service{
		executor: executor,
	}
}

func (s *service) Execute(in CodeExecRequest) (CodeExecResponse, error) {
	return s.executor.Execute(in)
}
