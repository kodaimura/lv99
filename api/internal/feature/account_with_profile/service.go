package account_with_profile

type Service interface {
	Get(in GetDto) ([]AccountWithProfile, error)
	GetOne(in GetOneDto) (AccountWithProfile, error)
}

type service struct {
	query Query
}

func NewService(query Query) Service {
	return &service{
		query: query,
	}
}

func (srv *service) Get(in GetDto) ([]AccountWithProfile, error) {
	return srv.query.GetWithProfile()
}

func (srv *service) GetOne(in GetOneDto) (AccountWithProfile, error) {
	return srv.query.GetOneWithProfile(in.Id)
}
