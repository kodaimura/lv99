package comment

type Service interface {
	GetWithProfile(in GetWithProfileDto) ([]CommentWithProfile, error)
}

type service struct {
	query Query
}

func NewService(query Query) Service {
	return &service{
		query: query,
	}
}

func (srv *service) GetWithProfile(in GetWithProfileDto) ([]CommentWithProfile, error) {
	return srv.query.GetWithProfile(in.AnswerId)
}
