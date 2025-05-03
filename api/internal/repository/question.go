package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Get(a *model.Question) ([]model.Question, error)
	GetOne(a *model.Question) (model.Question, error)

	Insert(a *model.Question) (model.Question, error)
	Update(a *model.Question) (model.Question, error)
	Delete(a *model.Question) error

	InsertTx(a *model.Question, tx *gorm.DB) (model.Question, error)
	UpdateTx(a *model.Question, tx *gorm.DB) (model.Question, error)
	DeleteTx(a *model.Question, tx *gorm.DB) error
}
