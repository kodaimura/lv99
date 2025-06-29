package account

type Service interface {
	GetWithProfile(in GetWithProfileDto) ([]AccountWithProfile, error)
	GetOneWithProfile(in GetOneWithProfileDto) (AccountWithProfile, error)
	GetAdminWithProfile() (AccountWithProfile, error)
}

type service struct {
	query Query
}

func NewService(query Query) Service {
	return &service{
		query: query,
	}
}

func (srv *service) GetWithProfile(in GetWithProfileDto) ([]AccountWithProfile, error) {
	return srv.query.GetWithProfile()
}

func (srv *service) GetOneWithProfile(in GetOneWithProfileDto) (AccountWithProfile, error) {
	return srv.query.GetOneWithProfile(in.Id)
}

func (srv *service) GetAdminWithProfile() (AccountWithProfile, error) {
	return srv.query.GetAdminWithProfile()
}
