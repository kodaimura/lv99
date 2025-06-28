package comment

import (
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Repository interface {
	Get(m *Comment, db *gorm.DB) ([]Comment, error)
	GetOne(m *Comment, db *gorm.DB) (Comment, error)
	GetAll(m *Comment, db *gorm.DB) ([]Comment, error)
	Insert(m *Comment, db *gorm.DB) (Comment, error)
	Update(m *Comment, db *gorm.DB) (Comment, error)
	Delete(m *Comment, db *gorm.DB) error

	RestoreOne(m *Comment, db *gorm.DB) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (rep *repository) Get(m *Comment, db *gorm.DB) ([]Comment, error) {
	var accounts []Comment
	err := db.Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *repository) GetOne(m *Comment, db *gorm.DB) (Comment, error) {
	var account Comment
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *repository) GetAll(m *Comment, db *gorm.DB) ([]Comment, error) {
	var accounts []Comment
	err := db.Unscoped().Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *repository) Insert(m *Comment, db *gorm.DB) (Comment, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *repository) Update(m *Comment, db *gorm.DB) (Comment, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *repository) Delete(m *Comment, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}

func (rep *repository) RestoreOne(m *Comment, db *gorm.DB) error {
	return db.Unscoped().
		Model(&Comment{}).
		Where("id = ?", m.Id).
		Update("deleted_at", nil).Error
}
