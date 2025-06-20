package question

import (
	"time"

	"gorm.io/gorm"
)

// ============================
// Question
// ============================

type QuestionResponse struct {
	Id        int            `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Answer    string         `json:"answer"`
	Level     int            `json:"level"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func ToQuestionResponse(m Question) QuestionResponse {
	return QuestionResponse(m)
}

func ToQuestionResponseList(models []Question) []QuestionResponse {
	res := make([]QuestionResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToQuestionResponse(m))
	}
	return res
}
