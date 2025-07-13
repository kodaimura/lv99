package question

import (
	questionModule "lv99/internal/module/question"

	"gorm.io/gorm"
)

type Usecase interface {
	Get(in GetDto) ([]questionModule.Question, error)
	GetAll(in GetAllDto) ([]questionModule.Question, error)
	GetOne(in GetOneDto) (questionModule.Question, error)
	CreateOne(in CreateOneDto) (questionModule.Question, error)
	UpdateOne(in UpdateOneDto) (questionModule.Question, error)
	DeleteOne(in DeleteOneDto) error
	RestoreOne(in RestoreOneDto) error
}

type usecase struct {
	db              *gorm.DB
	questionService questionModule.Service
}

func NewUsecase(
	db *gorm.DB,
	questionService questionModule.Service,
) Usecase {
	return &usecase{
		db:              db,
		questionService: questionService,
	}
}

func (uc *usecase) Get(in GetDto) ([]questionModule.Question, error) {
	return uc.questionService.Get(questionModule.Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, uc.db)
}

func (uc *usecase) GetAll(in GetAllDto) ([]questionModule.Question, error) {
	return uc.questionService.GetAll(questionModule.Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, uc.db)
}

func (uc *usecase) GetOne(in GetOneDto) (questionModule.Question, error) {
	return uc.questionService.GetOne(questionModule.Question{
		Id: in.Id,
	}, uc.db)
}

func (uc *usecase) CreateOne(in CreateOneDto) (questionModule.Question, error) {
	return uc.questionService.CreateOne(questionModule.Question{
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, uc.db)
}

func (uc *usecase) UpdateOne(in UpdateOneDto) (questionModule.Question, error) {
	return uc.questionService.UpdateOne(questionModule.Question{
		Id:      in.Id,
		Title:   in.Title,
		Content: in.Content,
		Answer:  in.Answer,
		Level:   in.Level,
	}, uc.db)
}

func (uc *usecase) DeleteOne(in DeleteOneDto) error {
	return uc.questionService.DeleteOne(questionModule.Question{
		Id: in.Id,
	}, uc.db)
}

func (uc *usecase) RestoreOne(in RestoreOneDto) error {
	return uc.questionService.RestoreOne(questionModule.Question{
		Id: in.Id,
	}, uc.db)
}
