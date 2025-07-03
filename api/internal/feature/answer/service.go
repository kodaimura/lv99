package answer

type Service interface {
	GetStatus(in GetStatusDto) ([]AnswerStatus, error)
	Search(in SearchDto) ([]AnswerSearch, error)
}
type service struct {
	query Query
}

func NewService(query Query) Service {
	return &service{
		query: query,
	}
}

func (srv *service) GetStatus(in GetStatusDto) ([]AnswerStatus, error) {
	return srv.query.GetStatus(
		in.AccountId,
	)
}

func (srv *service) Search(in SearchDto) ([]AnswerSearch, error) {
	return srv.query.Search(
		in.AccountId,
		in.QuestionId,
		in.Level,
		in.IsCorrect,
		in.CommentAccountId,
	)
}
