package response

import (
	"lv99/internal/model"
	"time"

	"gorm.io/gorm"
)

// ============================
// Comment
// ============================

type Comment struct {
	Id        int            `json:"id"`
	AnswerId  int            `json:"answer_id"`
	AccountId int            `json:"account_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromModelComment(m model.Comment) Comment {
	return Comment{
		Id:        m.Id,
		AnswerId:  m.AnswerId,
		AccountId: m.AccountId,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
	}
}

func FromModelCommentList(models []model.Comment) []Comment {
	res := make([]Comment, 0, len(models))
	for _, m := range models {
		res = append(res, FromModelComment(m))
	}
	return res
}
