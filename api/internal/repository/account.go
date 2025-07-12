package repository

import (
	"gorm.io/gorm"

	"lv99/internal/helper"
	"lv99/internal/model"
)

type AccountRepository interface {
	Get(m *model.Account, db *gorm.DB) ([]model.Account, error)
	GetOne(m *model.Account, db *gorm.DB) (model.Account, error)
	GetAll(m *model.Account, db *gorm.DB) ([]model.Account, error)
	Insert(m *model.Account, db *gorm.DB) (model.Account, error)
	Update(m *model.Account, db *gorm.DB) (model.Account, error)
	Delete(m *model.Account, db *gorm.DB) error
}

type accountRepository struct{}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

func (rep *accountRepository) Get(m *model.Account, db *gorm.DB) ([]model.Account, error) {
	var accounts []model.Account
	err := db.Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *accountRepository) GetOne(m *model.Account, db *gorm.DB) (model.Account, error) {
	var account model.Account
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *accountRepository) GetAll(m *model.Account, db *gorm.DB) ([]model.Account, error) {
	var accounts []model.Account
	err := db.Unscoped().Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *accountRepository) Insert(m *model.Account, db *gorm.DB) (model.Account, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *accountRepository) Update(m *model.Account, db *gorm.DB) (model.Account, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *accountRepository) Delete(m *model.Account, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}
