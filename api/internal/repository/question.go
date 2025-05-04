package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Get(q *model.Question) ([]model.Question, error)
	GetOne(q *model.Question) (model.Question, error)
	GetAll(q *model.Question) ([]model.Question, error)

	Insert(q *model.Question) (model.Question, error)
	Update(q *model.Question) (model.Question, error)
	Delete(q *model.Question) error

	InsertTx(q *model.Question, tx *gorm.DB) (model.Question, error)
	UpdateTx(q *model.Question, tx *gorm.DB) (model.Question, error)
	DeleteTx(q *model.Question, tx *gorm.DB) error

	RestoreOne(q *model.Question) error
}
