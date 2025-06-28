package answer

import (
	"lv99/internal/core"
	"lv99/internal/domain/executor"
	"lv99/internal/domain/question"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	Get(in GetDto, db *gorm.DB) ([]Answer, error)
	GetOne(in GetOneDto, db *gorm.DB) (Answer, error)
	CreateOne(in CreateOneDto, db *gorm.DB) (Answer, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (Answer, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
}

type service struct {
	repository      Repository
	questionService question.Service
	executorService executor.Service
}

func NewService(
	repository Repository,
	questionService question.Service,
	executorService executor.Service,
) Service {
	return &service{
		repository:      repository,
		questionService: questionService,
		executorService: executorService,
	}
}

func (srv *service) Get(in GetDto, db *gorm.DB) ([]Answer, error) {
	return srv.repository.Get(&Answer{
		QuestionId: in.QuestionId,
		AccountId:  in.AccountId,
	}, db)
}

func (srv *service) GetOne(in GetOneDto, db *gorm.DB) (Answer, error) {
	return srv.repository.GetOne(&Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}

func (srv *service) CreateOne(in CreateOneDto, db *gorm.DB) (Answer, error) {
	result, err := srv.executorService.Execute(executor.CodeExecRequest{
		CodeDef:  in.CodeDef,
		CodeCall: in.CodeCall,
	})
	if err != nil {
		return Answer{}, err
	}
	q, err := srv.questionService.GetOne(question.GetOneDto{Id: in.QuestionId}, db)
	if err != nil {
		core.Logger.Error(err.Error())
		return Answer{}, core.ErrBadRequest
	}

	isCorrect := false
	var correctAt *time.Time
	if result.Output == q.Answer {
		isCorrect = true
		now := time.Now()
		correctAt = &now
	}

	return srv.repository.Insert(&Answer{
		QuestionId: in.QuestionId,
		AccountId:  in.AccountId,
		CodeDef:    in.CodeDef,
		CodeCall:   in.CodeCall,
		CallOutput: result.Output,
		CallError:  result.Error,
		IsCorrect:  isCorrect,
		CorrectAt:  correctAt,
	}, db)
}

func (srv *service) UpdateOne(in UpdateOneDto, db *gorm.DB) (Answer, error) {
	ans, err := srv.repository.GetOne(&Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
	if err != nil {
		return Answer{}, err
	}

	result, err := srv.executorService.Execute(executor.CodeExecRequest{
		CodeDef:  in.CodeDef,
		CodeCall: in.CodeCall,
	})
	if err != nil {
		return Answer{}, err
	}
	q, err := srv.questionService.GetOne(question.GetOneDto{Id: in.QuestionId}, db)
	if err != nil {
		return Answer{}, err
	}

	ans.CodeDef = in.CodeDef
	ans.CodeCall = in.CodeCall
	ans.CallOutput = result.Output
	ans.CallError = result.Error
	ans.IsCorrect = result.Output == q.Answer

	if ans.IsCorrect && ans.CorrectAt.IsZero() {
		now := time.Now()
		ans.CorrectAt = &now
	}

	return srv.repository.Update(&ans, db)
}

func (srv *service) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return srv.repository.Delete(&Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}
