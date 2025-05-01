package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Get(a *model.Admin) ([]model.Admin, error)
	GetOne(a *model.Admin) (model.Admin, error)

	Insert(a *model.Admin) (model.Admin, error)
	Update(a *model.Admin) (model.Admin, error)
	Delete(a *model.Admin) error

	InsertTx(a *model.Admin, tx *gorm.DB) (model.Admin, error)
	UpdateTx(a *model.Admin, tx *gorm.DB) (model.Admin, error)
	DeleteTx(a *model.Admin, tx *gorm.DB) error
}
