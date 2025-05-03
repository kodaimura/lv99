package service

import (
	"lv99/internal/dto/input"
	"lv99/internal/model"
	"lv99/internal/repository"
)

type QuestionService interface {
	Get(in input.Question) ([]model.Question, error)
	GetOne(in input.Question) (model.Question, error)
	CreateOne(in input.Question) (model.Question, error)
	UpdateOne(in input.Question) (model.Question, error)
	DeleteOne(in input.Question) error
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
		QuestionTitle: in.QuestionTitle,
		QuestionContent: in.QuestionContent,
		QuestionAnswer: in.QuestionAnswer,
		QuestionLevel: in.QuestionLevel,
	})
}

func (srv *questionService) GetOne(in input.Question) (model.Question, error) {
	return srv.questionRepository.GetOne(&model.Question{QuestionId: in.QuestionId})
}

func (srv *questionService) CreateOne(in input.Question) (model.Question, error) {
	return srv.questionRepository.Insert(&model.Question{
		QuestionTitle:     in.QuestionTitle,
		QuestionContent: in.QuestionContent,
		QuestionAnswer: in.QuestionAnswer,
		QuestionLevel: in.QuestionLevel,
	})
}

func (srv *questionService) UpdateOne(in input.Question) (model.Question, error) {
	question, err := srv.questionRepository.GetOne(&model.Question{QuestionId: in.QuestionId})
	if err != nil {
		return model.Question{}, err
	}
	question.QuestionTitle = in.QuestionTitle
	question.QuestionContent = in.QuestionContent
	question.QuestionAnswer = in.QuestionAnswer
	question.QuestionLevel = in.QuestionLevel
	return srv.questionRepository.Update(&question)
}

func (srv *questionService) DeleteOne(in input.Question) error {
	return srv.questionRepository.Delete(&model.Question{QuestionId: in.QuestionId})
}
