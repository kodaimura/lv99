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
	question, err := srv.questionRepository.GetOne(&model.Question{QuestionId: in.QuestionId})
	if err != nil {
		return model.Answer{}, err
	}

	var isCorrect bool
	var correctAt time.Time
	var callOutput string
	var callError string
	if result.Output == question.QuestionAnswer {
		isCorrect = true
		correctAt = time.Now()
		callOutput = result.Output
	} else {
		callError = result.Error
	}

	return srv.answerRepository.Insert(&model.Answer{
		QuestionId: in.QuestionId,
		AccountId: in.AccountId,
		CodeDef: in.CodeDef,
		CodeCall: in.CodeCall,
		CallOutput: callOutput,
		CallError: callError,
		IsCorrect: isCorrect,
		CorrectAt: correctAt,
	})
}