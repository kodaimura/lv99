package response

import (
	"lv99/internal/model"
	"time"

	"gorm.io/gorm"
)

// ============================
// Answer
// ============================

type Answer struct {
	AnswerId   int            `json:"answer_id"`
	QuestionId int            `json:"question_id"`
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

func FromModelAnswer(m model.Answer) Answer {
	return Answer{
		AnswerId:   m.AnswerId,
		QuestionId: m.QuestionId,
		CodeDef:    m.CodeDef,
		CodeCall:   m.CodeCall,
		CallOutput: m.CallOutput,
		CallError:  m.CallError,
		IsCorrect:  m.IsCorrect,
		CorrectAt:  m.CorrectAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  m.DeletedAt,
	}
}

func FromModelAnswerList(models []model.Answer) []Answer {
	res := make([]Answer, 0, len(models))
	for _, m := range models {
		res = append(res, FromModelAnswer(m))
	}
	return res
}
