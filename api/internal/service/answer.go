package service

import (
	"time"

	"lv99/internal/dto/input"
	"lv99/internal/model"
	"lv99/internal/repository"
)

type AnswerService interface {
	Get(in input.Answer) ([]model.Answer, error)
	CreateOne(in input.Answer) (model.Answer, error)
	UpdateOne(in input.Answer) (model.Answer, error)
	DeleteOne(in input.Answer) error
}

type answerService struct {
	questionRepository repository.QuestionRepository
	answerRepository repository.AnswerRepository
	codeExecutor CodeExecutor
}

func NewAnswerService(
	questionRepository repository.QuestionRepository,
	answerRepository repository.AnswerRepository, 
	codeExecutor CodeExecutor,
) AnswerService {
	return &answerService{
		questionRepository: questionRepository,
		answerRepository: answerRepository,
		codeExecutor: codeExecutor,
	}
}

func (srv *answerService) Get(in input.Answer) ([]model.Answer, error) {
	return srv.answerRepository.Get(&model.Answer{
		QuestionId: in.QuestionId,
		AccountId: in.AccountId,
	})
}

func (srv *answerService) CreateOne(in input.Answer) (model.Answer, error) {
	result, err := srv.codeExecutor.Execute(CodeExecRequest{
		CodeDef: in.CodeDef,
		CodeCall: in.CodeCall,
	})
	if err != nil {
		return model.Answer{}, err
	}
	question, err := srv.questionRepository.GetOne(&model.Question{Id: in.QuestionId})
	if err != nil {
		return model.Answer{}, err
	}

	isCorrect := false
	var correctAt time.Time
	if result.Output == question.Answer {
		isCorrect = true
		correctAt = time.Now()
	}

	return srv.answerRepository.Insert(&model.Answer{
		QuestionId: in.QuestionId,
		AccountId: in.AccountId,
		CodeDef: in.CodeDef,
		CodeCall: in.CodeCall,
		CallOutput: result.Output,
		CallError: result.Error,
		IsCorrect: isCorrect,
		CorrectAt: correctAt,
	})
}

func (srv *answerService) UpdateOne(in input.Answer) (model.Answer, error) {
	answer, err := srv.answerRepository.GetOne(&model.Answer{
		AnswerId: in.AnswerId,
		QuestionId: in.QuestionId,
		AccountId: in.AccountId,
	})
	if err != nil {
		return model.Answer{}, err
	}

	result, err := srv.codeExecutor.Execute(CodeExecRequest{
		CodeDef: in.CodeDef,
		CodeCall: in.CodeCall,
	})
	if err != nil {
		return model.Answer{}, err
	}
	question, err := srv.questionRepository.GetOne(&model.Question{Id: in.QuestionId})
	if err != nil {
		return model.Answer{}, err
	}

	answer.CodeDef = in.CodeDef
	answer.CodeCall = in.CodeCall
	answer.CallOutput = result.Output
	answer.CallError = result.Error
	answer.IsCorrect = result.Output == question.Answer
 
	if answer.IsCorrect && answer.CorrectAt.IsZero() {
		answer.CorrectAt = time.Now()
	}

	return srv.answerRepository.Update(&answer)
}

func (srv *answerService) DeleteOne(in input.Answer) error {
	return srv.answerRepository.Delete(&model.Answer{
		AnswerId: in.AnswerId,
		QuestionId: in.QuestionId,
		AccountId: in.AccountId,
	})
}