package comment

import (
	"time"

	"gorm.io/gorm"
)

// ============================
// Comment
// ============================

type CommentResponse struct {
	Id        int            `json:"id"`
	AnswerId  int            `json:"answer_id"`
	AccountId int            `json:"account_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func ToCommentResponse(m Comment) CommentResponse {
	return CommentResponse(m)
}

func ToCommentResponseList(models []Comment) []CommentResponse {
	res := make([]CommentResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToCommentResponse(m))
	}
	return res
}
