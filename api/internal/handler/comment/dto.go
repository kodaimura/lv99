package comment

import (
	"time"

	"gorm.io/gorm"

	commentModule "lv99/internal/module/comment"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type CommentResponse struct {
	Id        int            `json:"id"`
	AnswerId  int            `json:"answer_id"`
	AccountId int            `json:"account_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func ToCommentResponse(m commentModule.Comment) CommentResponse {
	return CommentResponse(m)
}

func ToCommentResponseList(models []commentModule.Comment) []CommentResponse {
	res := make([]CommentResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToCommentResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type CommentUri struct {
	CommentId int `uri:"comment_id" binding:"required"`
}

type GetCommentsRequest struct {
	AnswerId int `form:"answer_id"`
}

type PostOneRequest struct {
	AnswerId int    `json:"answer_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type PutOneRequest struct {
	Content string `json:"content" binding:"required"`
}
