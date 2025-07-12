package model

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	Id         int            `db:"id" gorm:"column:id;primaryKey;autoIncrement"`
	QuestionId int            `db:"question_id" gorm:"column:question_id"`
	AccountId  int            `db:"account_id" gorm:"column:account_id"`
	CodeDef    string         `db:"code_def" gorm:"column:code_def"`
	CodeCall   string         `db:"code_call" gorm:"column:code_call"`
	CallOutput string         `db:"call_output" gorm:"column:call_output"`
	CallError  string         `db:"call_error" gorm:"column:call_error"`
	IsCorrect  bool           `db:"is_correct" gorm:"column:is_correct"`
	CorrectAt  *time.Time     `db:"correct_at" gorm:"column:correct_at"`
	CreatedAt  time.Time      `db:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time      `db:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `db:"deleted_at" gorm:"column:deleted_at;index"`
}
