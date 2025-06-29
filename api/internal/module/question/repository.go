package question

import (
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Repository interface {
	Get(m *Question, db *gorm.DB) ([]Question, error)
	GetOne(m *Question, db *gorm.DB) (Question, error)
	GetAll(m *Question, db *gorm.DB) ([]Question, error)
	Insert(m *Question, db *gorm.DB) (Question, error)
	Update(m *Question, db *gorm.DB) (Question, error)
	Delete(m *Question, db *gorm.DB) error

	RestoreOne(m *Question, db *gorm.DB) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (rep *repository) Get(m *Question, db *gorm.DB) ([]Question, error) {
	var accounts []Question
	err := db.Order("level ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *repository) GetOne(m *Question, db *gorm.DB) (Question, error) {
	var account Question
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *repository) GetAll(m *Question, db *gorm.DB) ([]Question, error) {
	var accounts []Question
	err := db.Unscoped().Order("level ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *repository) Insert(m *Question, db *gorm.DB) (Question, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *repository) Update(m *Question, db *gorm.DB) (Question, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *repository) Delete(m *Question, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}

func (rep *repository) RestoreOne(m *Question, db *gorm.DB) error {
	return db.Unscoped().
		Model(&Question{}).
		Where("id = ?", m.Id).
		Update("deleted_at", nil).Error
}
