package answer

import (
	"lv99/internal/core"
	"lv99/internal/module/executor"
	"time"

	answerModule "lv99/internal/module/answer"
	questionModule "lv99/internal/module/question"

	"gorm.io/gorm"
)

type Usecase interface {
	Get(in GetDto, db *gorm.DB) ([]answerModule.Answer, error)
	GetOne(in GetOneDto, db *gorm.DB) (answerModule.Answer, error)
	CreateOne(in CreateOneDto, db *gorm.DB) (answerModule.Answer, error)
	UpdateOne(in UpdateOneDto, db *gorm.DB) (answerModule.Answer, error)
	DeleteOne(in DeleteOneDto, db *gorm.DB) error
}

type usecase struct {
	answerService   answerModule.Service
	questionService questionModule.Service
	executorService executor.Service
}

func NewUsecase(
	answerService answerModule.Service,
	questionService questionModule.Service,
	executorService executor.Service,
) Usecase {
	return &usecase{
		answerService:   answerService,
		questionService: questionService,
		executorService: executorService,
	}
}

func (srv *usecase) Get(in GetDto, db *gorm.DB) ([]answerModule.Answer, error) {
	return srv.answerService.Get(answerModule.Answer{
		QuestionId: in.QuestionId,
		AccountId:  in.AccountId,
	}, db)
}

func (srv *usecase) GetOne(in GetOneDto, db *gorm.DB) (answerModule.Answer, error) {
	return srv.answerService.GetOne(answerModule.Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}

func (srv *usecase) CreateOne(in CreateOneDto, db *gorm.DB) (answerModule.Answer, error) {
	result, err := srv.executorService.Execute(executor.CodeExecRequest{
		CodeDef:  in.CodeDef,
		CodeCall: in.CodeCall,
	})
	if err != nil {
		return answerModule.Answer{}, err
	}
	q, err := srv.questionService.GetOne(questionModule.Question{Id: in.QuestionId}, db)
	if err != nil {
		core.Logger.Error(err.Error())
		return answerModule.Answer{}, core.ErrBadRequest
	}

	isCorrect := false
	var correctAt *time.Time
	if result.Output == q.Answer {
		isCorrect = true
		now := time.Now()
		correctAt = &now
	}

	return srv.answerService.CreateOne(answerModule.Answer{
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

func (srv *usecase) UpdateOne(in UpdateOneDto, db *gorm.DB) (answerModule.Answer, error) {
	ans, err := srv.answerService.GetOne(answerModule.Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
	if err != nil {
		return answerModule.Answer{}, err
	}

	result, err := srv.executorService.Execute(executor.CodeExecRequest{
		CodeDef:  in.CodeDef,
		CodeCall: in.CodeCall,
	})
	if err != nil {
		return answerModule.Answer{}, err
	}
	q, err := srv.questionService.GetOne(questionModule.Question{Id: ans.QuestionId}, db)
	if err != nil {
		return answerModule.Answer{}, err
	}

	ans.CodeDef = in.CodeDef
	ans.CodeCall = in.CodeCall
	ans.CallOutput = result.Output
	ans.CallError = result.Error
	ans.IsCorrect = result.Output == q.Answer

	if ans.IsCorrect && (ans.CorrectAt == nil || ans.CorrectAt.IsZero()) {
		now := time.Now()
		ans.CorrectAt = &now
	}

	return srv.answerService.UpdateOne(ans, db)
}

func (srv *usecase) DeleteOne(in DeleteOneDto, db *gorm.DB) error {
	return srv.answerService.DeleteOne(answerModule.Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, db)
}
