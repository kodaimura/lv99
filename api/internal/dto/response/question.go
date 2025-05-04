package response

import (
	"lv99/internal/model"
	"time"

	"gorm.io/gorm"
)

// ============================
// Question
// ============================

type Question struct {
	QuestionId      int            `json:"question_id"`
	QuestionTitle   string         `json:"question_title"`
	QuestionContent string         `json:"question_content"`
	QuestionAnswer  string         `json:"question_answer"`
	QuestionLevel   int            `json:"question_level"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

func FromModelQuestion(m model.Question) Question {
	return Question{
		QuestionId:      m.QuestionId,
		QuestionTitle:   m.QuestionTitle,
		QuestionContent: m.QuestionContent,
		QuestionAnswer:  m.QuestionAnswer,
		QuestionLevel:   m.QuestionLevel,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		DeletedAt:       m.DeletedAt,
	}
}

func FromModelQuestionList(models []model.Question) []Question {
	res := make([]Question, 0, len(models))
	for _, m := range models {
		res = append(res, FromModelQuestion(m))
	}
	return res
}
