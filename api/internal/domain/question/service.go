package question

import (
	"gorm.io/gorm"
)

type Service interface {
	Get(in GetDto, db *gorm.DB) ([]Question, error)
	GetAll(in GetAllDto, db *gorm.DB) ([]Question, error)
	GetOne(in GetOneDto, db *gorm.DB) (Question, error)
	CreateOne(in CreateOneDto, db *gorm.DB) (Question, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (Question, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
	RestoreOne(in RestoreOneDto, db *gorm.DB) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (srv *service) Get(in GetDto, db *gorm.DB) ([]Question, error) {
	return srv.repository.Get(&Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, db)
}

func (srv *service) GetAll(in GetAllDto, db *gorm.DB) ([]Question, error) {
	return srv.repository.GetAll(&Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, db)
}

func (srv *service) GetOne(in GetOneDto, db *gorm.DB) (Question, error) {
	return srv.repository.GetOne(&Question{Id: in.Id}, db)
}

func (srv *service) CreateOne(in CreateOneDto, db *gorm.DB) (Question, error) {
	return srv.repository.Insert(&Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, db)
}

func (srv *service) UpdateOne(in UpdateOneDto, db *gorm.DB) (Question, error) {
	question, err := srv.repository.GetOne(&Question{Id: in.Id}, db)
	if err != nil {
		return Question{}, err
	}
	question.Title = in.Title
	question.Content = in.Content
	question.Answer = in.Answer
	question.Level = in.Level
	return srv.repository.Update(&question, db)
}

func (srv *service) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return srv.repository.Delete(&Question{Id: in.Id}, db)
}

func (srv *service) RestoreOne(in RestoreOneDto, db *gorm.DB) error {
	return srv.repository.RestoreOne(&Question{Id: in.Id}, db)
}
