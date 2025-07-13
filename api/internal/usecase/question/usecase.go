package question

import (
	questionModule "lv99/internal/module/question"

	"gorm.io/gorm"
)

type Usecase interface {
	Get(in GetDto, db *gorm.DB) ([]questionModule.Question, error)
	GetAll(in GetAllDto, db *gorm.DB) ([]questionModule.Question, error)
	GetOne(in GetOneDto, db *gorm.DB) (questionModule.Question, error)
	CreateOne(in CreateOneDto, db *gorm.DB) (questionModule.Question, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (questionModule.Question, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
	RestoreOne(in RestoreOneDto, db *gorm.DB) error
}

type usecase struct {
	service questionModule.Service
}

func NewUsecase(service questionModule.Service) Usecase {
	return &usecase{
		service: service,
	}
}

func (srv *usecase) Get(in GetDto, db *gorm.DB) ([]questionModule.Question, error) {
	return srv.service.Get(questionModule.Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, db)
}

func (srv *usecase) GetAll(in GetAllDto, db *gorm.DB) ([]questionModule.Question, error) {
	return srv.service.GetAll(questionModule.Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, db)
}

func (srv *usecase) GetOne(in GetOneDto, db *gorm.DB) (questionModule.Question, error) {
	return srv.service.GetOne(questionModule.Question{Id: in.Id}, db)
}

func (srv *usecase) CreateOne(in CreateOneDto, db *gorm.DB) (questionModule.Question, error) {
	return srv.service.CreateOne(questionModule.Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, db)
}

func (srv *usecase) UpdateOne(in UpdateOneDto, db *gorm.DB) (questionModule.Question, error) {
	return srv.service.UpdateOne(questionModule.Question{
		Id:      in.Id,
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, db)
}

func (srv *usecase) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return srv.service.DeleteOne(questionModule.Question{Id: in.Id}, db)
}

func (srv *usecase) RestoreOne(in RestoreOneDto, db *gorm.DB) error {
	return srv.service.RestoreOne(questionModule.Question{Id: in.Id}, db)
}
