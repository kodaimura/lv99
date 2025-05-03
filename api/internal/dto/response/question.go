package response

import (
	"lv99/internal/model"
)

// ============================
// Question
// ============================

type Question struct {
	QuestionId      int     `json:"question_id"`
	QuestionTitle   string  `json:"question_title"`
	QuestionContent string  `json:"question_content"`
	QuestionAnswer  string  `json:"question_answer"`
	QuestionLevel   int     `json:"question_level"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
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
	}
}

func FromModelQuestionList(models []model.Question) []Question {
	res := make([]Question, 0, len(models))
	for _, m := range models {
		res = append(res, FromModelQuestion(m))
	}
	return res
}