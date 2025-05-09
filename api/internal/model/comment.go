package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	CommentId      int            `db:"comment_id" json:"comment_id" gorm:"primaryKey;autoIncrement"`
	AnswerId       int            `db:"answer_id" json:"answer_id"`
	AccountId      int            `db:"account_id" json:"account_id"`
	CommentContent string         `db:"comment_content" json:"comment_content"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time      `db:"updated_at" json:"updated_at" gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `db:"deleted_at" json:"deleted_at" gorm:"index"`
}
