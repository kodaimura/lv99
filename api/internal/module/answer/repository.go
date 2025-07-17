package answer

import (
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Repository interface {
	Get(m *Answer, db *gorm.DB) ([]Answer, error)
	GetOne(m *Answer, db *gorm.DB) (Answer, error)
	GetAll(m *Answer, db *gorm.DB) ([]Answer, error)
	Insert(m *Answer, db *gorm.DB) (Answer, error)
	Update(m *Answer, db *gorm.DB) (Answer, error)
	Delete(m *Answer, db *gorm.DB) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (rep *repository) Get(m *Answer, db *gorm.DB) ([]Answer, error) {
	var accounts []Answer
	err := db.Order("id DESC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *repository) GetOne(m *Answer, db *gorm.DB) (Answer, error) {
	var account Answer
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *repository) GetAll(m *Answer, db *gorm.DB) ([]Answer, error) {
	var accounts []Answer
	err := db.Unscoped().Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *repository) Insert(m *Answer, db *gorm.DB) (Answer, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *repository) Update(m *Answer, db *gorm.DB) (Answer, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *repository) Delete(m *Answer, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}
