package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type QuestionRepository interface {
	Get(m *model.Question) ([]model.Question, error)
	GetOne(m *model.Question) (model.Question, error)
	GetAll(m *model.Question) ([]model.Question, error)

	Insert(m *model.Question) (model.Question, error)
	Update(m *model.Question) (model.Question, error)
	Delete(m *model.Question) error

	InsertTx(m *model.Question, tx *gorm.DB) (model.Question, error)
	UpdateTx(m *model.Question, tx *gorm.DB) (model.Question, error)
	DeleteTx(m *model.Question, tx *gorm.DB) error

	RestoreOne(m *model.Question) error
}
