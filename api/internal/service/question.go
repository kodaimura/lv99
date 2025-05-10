package service

import (
	"lv99/internal/dto/input"
	"lv99/internal/model"
	"lv99/internal/repository"
)

type QuestionService interface {
	Get(in input.Question) ([]model.Question, error)
	GetAll(in input.Question) ([]model.Question, error)
	GetOne(in input.Question) (model.Question, error)
	CreateOne(in input.Question) (model.Question, error)
	UpdateOne(in input.Question) (model.Question, error)
	DeleteOne(in input.Question) error
	RestoreOne(in input.Question) error
}

type questionService struct {
	questionRepository repository.QuestionRepository
}

func NewQuestionService(questionRepository repository.QuestionRepository) QuestionService {
	return &questionService{
		questionRepository: questionRepository,
	}
}

func (srv *questionService) Get(in input.Question) ([]model.Question, error) {
	return srv.questionRepository.Get(&model.Question{
		Title: in.Title,
		Content: in.Content,
		Answer: in.Answer,
		Level: in.Level,
	})
}

func (srv *questionService) GetAll(in input.Question) ([]model.Question, error) {
	return srv.questionRepository.GetAll(&model.Question{
		Title: in.Title,
		Content: in.Content,
		Answer: in.Answer,
		Level: in.Level,
	})
}

func (srv *questionService) GetOne(in input.Question) (model.Question, error) {
	return srv.questionRepository.GetOne(&model.Question{Id: in.Id})
}

func (srv *questionService) CreateOne(in input.Question) (model.Question, error) {
	return srv.questionRepository.Insert(&model.Question{
		Title:     in.Title,
		Content: in.Content,
		Answer: in.Answer,
		Level: in.Level,
	})
}

func (srv *questionService) UpdateOne(in input.Question) (model.Question, error) {
	question, err := srv.questionRepository.GetOne(&model.Question{Id: in.Id})
	if err != nil {
		return model.Question{}, err
	}
	question.Title = in.Title
	question.Content = in.Content
	question.Answer = in.Answer
	question.Level = in.Level
	return srv.questionRepository.Update(&question)
}

func (srv *questionService) DeleteOne(in input.Question) error {
	return srv.questionRepository.Delete(&model.Question{Id: in.Id})
}

func (srv *questionService) RestoreOne(in input.Question) error {
	return srv.questionRepository.RestoreOne(&model.Question{Id: in.Id})
}
