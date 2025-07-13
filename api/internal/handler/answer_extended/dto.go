package answer_extended

import (
	usecase "lv99/internal/usecase/answer_extended"
	"time"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type AnswerSearchResponse struct {
	AnswerId           int        `json:"answer_id"`
	CodeDef            string     `json:"code_def"`
	CodeCall           string     `json:"code_call"`
	IsCorrect          bool       `json:"is_correct"`
	CorrectAt          *time.Time `json:"correct_at"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	QuestionId         int        `json:"question_id"`
	QuestionTitle      string     `json:"question_title"`
	QuestionLevel      int        `json:"question_level"`
	AccountId          int        `json:"account_id"`
	AccountName        string     `json:"account_name"`
	CommentCount       int        `json:"comment_count"`
	CommentAccountId   *int       `json:"comment_account_id"`
	CommentAccountName *string    `json:"comment_account_name"`
	CommentAt          *time.Time `json:"comment_at"`
}

func ToAnswerSearchResponse(m usecase.AnswerSearch) AnswerSearchResponse {
	return AnswerSearchResponse(m)
}

func ToAnswerSearchResponseList(models []usecase.AnswerSearch) []AnswerSearchResponse {
	res := make([]AnswerSearchResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAnswerSearchResponse(m))
	}
	return res
}

type AnswerStatusResponse struct {
	QuestionId   int        `json:"question_id"`
	IsCorrect    bool       `json:"is_correct"`
	CorrectCount int        `json:"correct_count"`
	CorrectAt    *time.Time `json:"correct_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func ToAnswerStatusResponse(m usecase.AnswerStatus) AnswerStatusResponse {
	return AnswerStatusResponse(m)
}

func ToAnswerStatusResponseList(models []usecase.AnswerStatus) []AnswerStatusResponse {
	res := make([]AnswerStatusResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAnswerStatusResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type SearchRequest struct {
	QuestionId       int   `form:"question_id"`
	AccountId        int   `form:"account_id"`
	Level            int   `form:"level"`
	IsCorrect        *bool `form:"is_correct"`
	CommentAccountId int   `form:"comment_account_id"`
}

type GetStatusRequest struct {
	AccountId int `form:"account_id"`
}
