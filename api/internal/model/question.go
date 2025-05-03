package model

import (
	"gorm.io/gorm"
)

type Question struct {
	QuestionId      int     `db:"question_id" json:"question_id" gorm:"primaryKey;autoIncrement"`
	QuestionTitle   string  `db:"question_title" json:"question_title"`
	QuestionContent string  `db:"question_content" json:"question_content"`
	QuestionAnswer  string  `db:"question_answer" json:"question_answer"`
	QuestionLevel   int     `db:"question_level" json:"question_level"`
	CreatedAt       string  `db:"created_at" json:"created_at" gorm:"-"`
	UpdatedAt       string  `db:"updated_at" json:"updated_at" gorm:"-"`
	DeletedAt       gorm.DeletedAt `db:"deleted_at" json:"deleted_at" gorm:"index"`
}