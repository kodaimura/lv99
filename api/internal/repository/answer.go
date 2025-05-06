package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type AnswerRepository interface {
	Get(a *model.Answer) ([]model.Answer, error)
	GetOne(a *model.Answer) (model.Answer, error)
	GetAll(a *model.Answer) ([]model.Answer, error)

	Insert(a *model.Answer) (model.Answer, error)
	Update(a *model.Answer) (model.Answer, error)
	Delete(a *model.Answer) error

	InsertTx(a *model.Answer, tx *gorm.DB) (model.Answer, error)
	UpdateTx(a *model.Answer, tx *gorm.DB) (model.Answer, error)
	DeleteTx(a *model.Answer, tx *gorm.DB) error
}
