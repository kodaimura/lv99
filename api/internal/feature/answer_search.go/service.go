package answer_search

type Service interface {
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
func (srv *service) Search(in SearchDto) ([]AnswerSearch, error) {
	return srv.query.Search(
		in.QuestionId,
		in.AccountId,
		in.IsCorrect,
		in.CommentAccountId,
	)
}
