package repository

import (
	"gorm.io/gorm"

	"lv99/internal/helper"
	"lv99/internal/model"
)

type AnswerRepository interface {
	Get(m *model.Answer, db *gorm.DB) ([]model.Answer, error)
	GetOne(m *model.Answer, db *gorm.DB) (model.Answer, error)
	GetAll(m *model.Answer, db *gorm.DB) ([]model.Answer, error)
	Insert(m *model.Answer, db *gorm.DB) (model.Answer, error)
	Update(m *model.Answer, db *gorm.DB) (model.Answer, error)
	Delete(m *model.Answer, db *gorm.DB) error
}

type answerAnswerRepository struct{}

func NewAnswerRepository() AnswerRepository {
	return &answerAnswerRepository{}
}

func (rep *answerAnswerRepository) Get(m *model.Answer, db *gorm.DB) ([]model.Answer, error) {
	var accounts []model.Answer
	err := db.Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *answerAnswerRepository) GetOne(m *model.Answer, db *gorm.DB) (model.Answer, error) {
	var account model.Answer
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *answerAnswerRepository) GetAll(m *model.Answer, db *gorm.DB) ([]model.Answer, error) {
	var accounts []model.Answer
	err := db.Unscoped().Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *answerAnswerRepository) Insert(m *model.Answer, db *gorm.DB) (model.Answer, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *answerAnswerRepository) Update(m *model.Answer, db *gorm.DB) (model.Answer, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *answerAnswerRepository) Delete(m *model.Answer, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}
