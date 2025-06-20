package answer

import (
	"time"

	"gorm.io/gorm"
)

// ============================
// Answer
// ============================

type AnswerResponse struct {
	Id         int            `json:"id"`
	QuestionId int            `json:"question_id"`
	AccountId  int            `json:"account_id"`
	CodeDef    string         `json:"code_def"`
	CodeCall   string         `json:"code_call"`
	CallOutput string         `json:"call_output"`
	CallError  string         `json:"call_error"`
	IsCorrect  bool           `json:"is_correct"`
	CorrectAt  time.Time      `json:"correct_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func ToAnswerResponse(m Answer) AnswerResponse {
	return AnswerResponse{
		Id:         m.Id,
		QuestionId: m.QuestionId,
		AccountId:  m.AccountId,
		CodeDef:    m.CodeDef,
		CodeCall:   m.CodeCall,
		CallOutput: m.CallOutput,
		CallError:  m.CallError,
		IsCorrect:  m.IsCorrect,
		CorrectAt:  m.CorrectAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func ToAnswerResponseList(models []Answer) []AnswerResponse {
	res := make([]AnswerResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAnswerResponse(m))
	}
	return res
}
