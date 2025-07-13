package question

import (
	"time"

	"gorm.io/gorm"

	questionModule "lv99/internal/module/question"
)

// -----------------------------
// DTO（Response）
// -----------------------------

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

func ToQuestionResponse(m questionModule.Question) QuestionResponse {
	return QuestionResponse(m)
}

func ToQuestionResponseList(models []questionModule.Question) []QuestionResponse {
	res := make([]QuestionResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToQuestionResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type QuestionUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
}

type PostOneRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Level   int    `json:"level" binding:"required,min=1"`
}

type PutOneRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Level   int    `json:"level" binding:"required,min=1"`
}
