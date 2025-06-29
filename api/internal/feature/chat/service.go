package chat

type Service interface {
	GetUnreadCount(in GetUnreadCountDto) ([]UnreadCount, error)
}

type service struct {
	query Query
}

func NewService(query Query) Service {
	return &service{
		query: query,
	}
}

func (srv *service) GetUnreadCount(in GetUnreadCountDto) ([]UnreadCount, error) {
	return srv.query.GetUnreadCount(in.ToId)
}
