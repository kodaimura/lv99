package answer

import (
	"lv99/internal/core"
	"time"

	answerModule "lv99/internal/module/answer"
	questionModule "lv99/internal/module/question"

	"gorm.io/gorm"
)

type Usecase interface {
	Get(in GetDto) ([]answerModule.Answer, error)
	GetOne(in GetOneDto) (answerModule.Answer, error)
	CreateOne(in CreateOneDto) (answerModule.Answer, error)
	UpdateOne(in UpdateOneDto) (answerModule.Answer, error)
	DeleteOne(in DeleteOneDto) error
}

type usecase struct {
	db              *gorm.DB
	answerService   answerModule.Service
	questionService questionModule.Service
}

func NewUsecase(
	db *gorm.DB,
	answerService answerModule.Service,
	questionService questionModule.Service,
) Usecase {
	return &usecase{
		db:              db,
		answerService:   answerService,
		questionService: questionService,
	}
}

func (uc *usecase) Get(in GetDto) ([]answerModule.Answer, error) {
	return uc.answerService.Get(answerModule.Answer{
		QuestionId: in.QuestionId,
		AccountId:  in.AccountId,
	}, uc.db)
}

func (uc *usecase) GetOne(in GetOneDto) (answerModule.Answer, error) {
	return uc.answerService.GetOne(answerModule.Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, uc.db)
}

func (uc *usecase) CreateOne(in CreateOneDto) (answerModule.Answer, error) {
	result, err := core.CodeExecutor.Execute(core.CodeExecRequest{
		CodeDef:  in.CodeDef,
		CodeCall: in.CodeCall,
	})
	if err != nil {
		return answerModule.Answer{}, err
	}

	q, err := uc.questionService.GetOne(questionModule.Question{
		Id: in.QuestionId,
	}, uc.db)
	if err != nil {
		return answerModule.Answer{}, core.ErrBadRequest
	}

	isCorrect := false
	var correctAt *time.Time
	if result.Output == q.Answer {
		isCorrect = true
		now := time.Now()
		correctAt = &now
	}

	return uc.answerService.CreateOne(answerModule.Answer{
		QuestionId: in.QuestionId,
		AccountId:  in.AccountId,
		CodeDef:    in.CodeDef,
		CodeCall:   in.CodeCall,
		CallOutput: result.Output,
		CallError:  result.Error,
		IsCorrect:  isCorrect,
		CorrectAt:  correctAt,
	}, uc.db)
}

func (uc *usecase) UpdateOne(in UpdateOneDto) (answerModule.Answer, error) {
	ans, err := uc.answerService.GetOne(answerModule.Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, uc.db)
	if err != nil {
		return answerModule.Answer{}, err
	}

	result, err := core.CodeExecutor.Execute(core.CodeExecRequest{
		CodeDef:  in.CodeDef,
		CodeCall: in.CodeCall,
	})
	if err != nil {
		return answerModule.Answer{}, err
	}
	q, err := uc.questionService.GetOne(questionModule.Question{
		Id: ans.QuestionId,
	}, uc.db)
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

	return uc.answerService.UpdateOne(ans, uc.db)
}

func (uc *usecase) DeleteOne(in DeleteOneDto) error {
	return uc.answerService.DeleteOne(answerModule.Answer{
		Id:        in.Id,
		AccountId: in.AccountId,
	}, uc.db)
}
