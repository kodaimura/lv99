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
	Id        int            `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Answer    string         `json:"answer"`
	Level     int            `json:"level"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromModelQuestion(m model.Question) Question {
	return Question{
		Id:        m.Id,
		Title:     m.Title,
		Content:   m.Content,
		Answer:    m.Answer,
		Level:     m.Level,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
	}
}

func FromModelQuestionList(models []model.Question) []Question {
	res := make([]Question, 0, len(models))
	for _, m := range models {
		res = append(res, FromModelQuestion(m))
	}
	return res
}
