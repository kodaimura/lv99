package answer

import (
	"time"
)

// ============================
// AnswerSearchResponse
// ============================

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

func ToAnswerSearchResponse(m AnswerSearch) AnswerSearchResponse {
	return AnswerSearchResponse(m)
}

func ToAnswerSearchResponseList(models []AnswerSearch) []AnswerSearchResponse {
	res := make([]AnswerSearchResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAnswerSearchResponse(m))
	}
	return res
}

// ============================
// AnswerStatusResponse
// ============================

type AnswerStatusResponse struct {
	QuestionId   int        `json:"question_id"`
	IsCorrect    bool       `json:"is_correct"`
	CorrectCount int        `json:"correct_count"`
	CorrectAt    *time.Time `json:"correct_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func ToAnswerStatusResponse(m AnswerStatus) AnswerStatusResponse {
	return AnswerStatusResponse(m)
}

func ToAnswerStatusResponseList(models []AnswerStatus) []AnswerStatusResponse {
	res := make([]AnswerStatusResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAnswerStatusResponse(m))
	}
	return res
}
