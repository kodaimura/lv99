package chat

import (
	"gorm.io/gorm"
)

type Service interface {
	Get(in GetDto, db *gorm.DB) ([]Chat, error)
	CreateOne(in CreateOneDto, db *gorm.DB) (Chat, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (Chat, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
}

type service struct {
	repository Repository
	query      Query
}

func NewService(
	repository Repository,
	query Query,
) Service {
	return &service{
		repository: repository,
		query:      query,
	}
}

func (srv *service) Get(in GetDto, db *gorm.DB) ([]Chat, error) {
	return srv.query.Get(in.FromId, in.ToId, in.Before, in.Limit)
}

func (srv *service) CreateOne(in CreateOneDto, db *gorm.DB) (Chat, error) {
	return srv.repository.Insert(&Chat{
		FromId:  in.FromId,
		ToId:    in.ToId,
		Message: in.Message,
	}, db)
}

func (srv *service) UpdateOne(in UpdateOneDto, db *gorm.DB) (Chat, error) {
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

func (srv *service) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return srv.repository.Delete(&Chat{Id: in.Id}, db)
}
