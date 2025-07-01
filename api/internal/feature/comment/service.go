package comment

type Service interface {
	GetWithProfile(in GetWithProfileDto) ([]CommentWithProfile, error)
	GetCount(in GetCountDto) ([]CommentCount, error)
	GetCountForAdmin(in GetCountDto) ([]CommentCount, error)
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

func (srv *service) GetCount(in GetCountDto) ([]CommentCount, error) {
	return srv.query.GetCount(in.AccountId, in.Since)
}

func (srv *service) GetCountForAdmin(in GetCountDto) ([]CommentCount, error) {
	return srv.query.GetCountForAdmin(in.AccountId, in.Since)
}
