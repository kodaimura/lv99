package chat

import (
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Service interface {
	Get(in Chat, db *gorm.DB) ([]Chat, error)
	GetOne(in Chat, db *gorm.DB) (Chat, error)
	CreateOne(in Chat, db *gorm.DB) (Chat, error)
	UpdateOne(in Chat, db *gorm.DB) (Chat, error)
	DeleteOne(in Chat, db *gorm.DB) error
	Read(in Chat, db *gorm.DB) error
	Paginate(accounId1 int, accountId2 int, before time.Time, limit int, db *sqlx.DB) ([]Chat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (srv *service) Get(in Chat, db *gorm.DB) ([]Chat, error) {
	return srv.repository.Get(&Chat{
		FromId: in.FromId,
		ToId:   in.ToId,
	}, db)
}

func (srv *service) GetOne(in Chat, db *gorm.DB) (Chat, error) {
	return srv.repository.GetOne(&Chat{
		Id: in.Id,
	}, db)
}

func (srv *service) CreateOne(in Chat, db *gorm.DB) (Chat, error) {
	return srv.repository.Insert(&Chat{
		FromId:  in.FromId,
		ToId:    in.ToId,
		Message: in.Message,
	}, db)
}

func (srv *service) UpdateOne(in Chat, db *gorm.DB) (Chat, error) {
	chat, err := srv.repository.GetOne(&Chat{Id: in.Id}, db)
	if err != nil {
		return Chat{}, err
	}

	chat.Message = in.Message
	if in.IsRead {
		chat.IsRead = in.IsRead
	}
	return srv.repository.Update(&chat, db)
}

func (srv *service) DeleteOne(in Chat, db *gorm.DB) error {
	return srv.repository.Delete(&Chat{Id: in.Id}, db)
}

func (srv *service) Read(in Chat, db *gorm.DB) error {
	return srv.repository.Read(&Chat{FromId: in.FromId, ToId: in.ToId}, db)
}

func (srv *service) Paginate(accounId1 int, accountId2 int, before time.Time, limit int, db *sqlx.DB) ([]Chat, error) {
	return srv.repository.Paginate(accounId1, accountId2, before, limit, db)
}
