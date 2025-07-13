package comment

import (
	"gorm.io/gorm"
)

type Service interface {
	Get(in Comment, db *gorm.DB) ([]Comment, error)
	GetOne(in Comment, db *gorm.DB) (Comment, error)
	CreateOne(in Comment, db *gorm.DB) (Comment, error)
	UpdateOne(in Comment, db *gorm.DB) (Comment, error)
	DeleteOne(in Comment, db *gorm.DB) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (srv *service) Get(in Comment, db *gorm.DB) ([]Comment, error) {
	return srv.repository.Get(&Comment{
		AnswerId:  in.AnswerId,
		AccountId: in.AccountId,
	}, db)
}

func (srv *service) GetOne(in Comment, db *gorm.DB) (Comment, error) {
	return srv.repository.GetOne(&Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}

func (srv *service) CreateOne(in Comment, db *gorm.DB) (Comment, error) {
	return srv.repository.Insert(&Comment{
		AnswerId:  in.AnswerId,
		AccountId: in.AccountId,
		Content:   in.Content,
	}, db)
}

func (srv *service) UpdateOne(in Comment, db *gorm.DB) (Comment, error) {
	comment, err := srv.repository.GetOne(&Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
	if err != nil {
		return Comment{}, err
	}

	comment.Content = in.Content
	return srv.repository.Update(&comment, db)
}

func (srv *service) DeleteOne(in Comment, db *gorm.DB) error {
	return srv.repository.Delete(&Comment{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}
