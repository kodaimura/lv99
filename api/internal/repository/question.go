package repository

import (
	"gorm.io/gorm"

	"lv99/internal/helper"
	"lv99/internal/model"
)

type QuestionRepository interface {
	Get(m *model.Question, db *gorm.DB) ([]model.Question, error)
	GetOne(m *model.Question, db *gorm.DB) (model.Question, error)
	GetAll(m *model.Question, db *gorm.DB) ([]model.Question, error)
	Insert(m *model.Question, db *gorm.DB) (model.Question, error)
	Update(m *model.Question, db *gorm.DB) (model.Question, error)
	Delete(m *model.Question, db *gorm.DB) error

	RestoreOne(m *model.Question, db *gorm.DB) error
}

type questionRepository struct{}

func NewQuestionRepository() QuestionRepository {
	return &questionRepository{}
}

func (rep *questionRepository) Get(m *model.Question, db *gorm.DB) ([]model.Question, error) {
	var accounts []model.Question
	err := db.Order("level ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *questionRepository) GetOne(m *model.Question, db *gorm.DB) (model.Question, error) {
	var account model.Question
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *questionRepository) GetAll(m *model.Question, db *gorm.DB) ([]model.Question, error) {
	var accounts []model.Question
	err := db.Unscoped().Order("level ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *questionRepository) Insert(m *model.Question, db *gorm.DB) (model.Question, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *questionRepository) Update(m *model.Question, db *gorm.DB) (model.Question, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *questionRepository) Delete(m *model.Question, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}

func (rep *questionRepository) RestoreOne(m *model.Question, db *gorm.DB) error {
	return db.Unscoped().
		Model(&model.Question{}).
		Where("id = ?", m.Id).
		Update("deleted_at", nil).Error
}
