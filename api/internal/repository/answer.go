package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Get(m *model.Answer) ([]model.Answer, error)
	GetOne(m *model.Answer) (model.Answer, error)
	GetAll(m *model.Answer) ([]model.Answer, error)

	Insert(m *model.Answer) (model.Answer, error)
	Update(m *model.Answer) (model.Answer, error)
	Delete(m *model.Answer) error

	InsertTx(m *model.Answer, tx *gorm.DB) (model.Answer, error)
	UpdateTx(m *model.Answer, tx *gorm.DB) (model.Answer, error)
	DeleteTx(m *model.Answer, tx *gorm.DB) error
}
