package model

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	AnswerId   int            `db:"answer_id" json:"answer_id" gorm:"primaryKey;autoIncrement"`
	QuestionId int            `db:"question_id" json:"question_id"`
	CodeDef    string         `db:"code_def" json:"code_def"`
	CodeCall   string         `db:"code_call" json:"code_call"`
	CallOutput string         `db:"call_output" json:"call_output"`
	CallError  string         `db:"call_error" json:"call_error"`
	IsCorrect  bool           `db:"is_correct" json:"is_correct"`
	CorrectAt  time.Time      `db:"correct_at" json:"correct_at"`
	CreatedAt  time.Time      `db:"created_at" json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time      `db:"updated_at" json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `db:"deleted_at" json:"deleted_at" gorm:"index"`
}
