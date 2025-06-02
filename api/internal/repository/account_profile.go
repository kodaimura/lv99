package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type AccountProfileRepository interface {
	Get(m *model.AccountProfile) ([]model.AccountProfile, error)
	GetOne(m *model.AccountProfile) (model.AccountProfile, error)
	GetAll(m *model.AccountProfile) ([]model.AccountProfile, error)

	Insert(m *model.AccountProfile) (model.AccountProfile, error)
	Update(m *model.AccountProfile) (model.AccountProfile, error)
	Delete(m *model.AccountProfile) error

	InsertTx(m *model.AccountProfile, tx *gorm.DB) (model.AccountProfile, error)
	UpdateTx(m *model.AccountProfile, tx *gorm.DB) (model.AccountProfile, error)
	DeleteTx(m *model.AccountProfile, tx *gorm.DB) error
}
