package comment_extended

import (
	usecase "lv99/internal/usecase/comment_extended"
	"time"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type CommentWithProfileResponse struct {
	Id          int       `json:"id"`
	AnswerId    int       `json:"answer_id"`
	AccountId   int       `json:"account_id"`
	DisplayName string    `json:"display_name"`
	AvatarURL   string    `json:"avatar_url"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToCommentWithProfileResponse(m usecase.CommentWithProfile) CommentWithProfileResponse {
	return CommentWithProfileResponse(m)
}

func ToCommentWithProfileResponseList(models []usecase.CommentWithProfile) []CommentWithProfileResponse {
	res := make([]CommentWithProfileResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToCommentWithProfileResponse(m))
	}
	return res
}

type CommentCountResponse struct {
	QuestionId    int       `json:"question_id"`
	QuestionTitle string    `json:"question_title"`
	QuestionLevel int       `json:"question_level"`
	AnswerId      int       `json:"answer_id"`
	CommentCount  int       `json:"comment_count"`
	CreatedAt     time.Time `json:"created_at"`
}

func ToCommentCountResponse(m usecase.CommentCount) CommentCountResponse {
	return CommentCountResponse(m)
}

func ToCommentCountResponseList(models []usecase.CommentCount) []CommentCountResponse {
	res := make([]CommentCountResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToCommentCountResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type GetWithProfileRequest struct {
	AnswerId int `form:"answer_id"`
}

type GetCountRequest struct {
	Since *time.Time `form:"since" time_format:"2006-01-02"`
}
