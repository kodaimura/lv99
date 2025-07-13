package answer

import (
	"time"

	"gorm.io/gorm"

	answerModule "lv99/internal/module/answer"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type AnswerResponse struct {
	Id         int            `json:"id"`
	QuestionId int            `json:"question_id"`
	AccountId  int            `json:"account_id"`
	CodeDef    string         `json:"code_def"`
	CodeCall   string         `json:"code_call"`
	CallOutput string         `json:"call_output"`
	CallError  string         `json:"call_error"`
	IsCorrect  bool           `json:"is_correct"`
	CorrectAt  *time.Time     `json:"correct_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func ToAnswerResponse(m answerModule.Answer) AnswerResponse {
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

func ToAnswerResponseList(models []answerModule.Answer) []AnswerResponse {
	res := make([]AnswerResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAnswerResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type AnswerUri struct {
	AnswerId int `uri:"answer_id" binding:"required"`
}

type QuestionAnswerUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
	AnswerId   int `uri:"answer_id" binding:"required"`
}

type GetAnswersRequest struct {
	QuestionId int `form:"question_id"`
}

type PostAnswersRequest struct {
	QuestionId int    `json:"question_id" binding:"required"`
	CodeDef    string `json:"code_def" binding:"required"`
	CodeCall   string `json:"code_call" binding:"required"`
}

type PutAnswerRequest struct {
	CodeDef  string `json:"code_def" binding:"required"`
	CodeCall string `json:"code_call" binding:"required"`
}

type AdminGetAnswersRequest struct {
	AccountId  int `form:"account_id"`
	QuestionId int `form:"question_id"`
}
