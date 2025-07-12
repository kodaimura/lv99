package repository

import (
	"gorm.io/gorm"

	"lv99/internal/helper"
	"lv99/internal/model"
)

type CommentRepository interface {
	Get(m *model.Comment, db *gorm.DB) ([]model.Comment, error)
	GetOne(m *model.Comment, db *gorm.DB) (model.Comment, error)
	GetAll(m *model.Comment, db *gorm.DB) ([]model.Comment, error)
	Insert(m *model.Comment, db *gorm.DB) (model.Comment, error)
	Update(m *model.Comment, db *gorm.DB) (model.Comment, error)
	Delete(m *model.Comment, db *gorm.DB) error

	RestoreOne(m *model.Comment, db *gorm.DB) error
}

type commentCommentRepository struct{}

func NewCommentRepository() CommentRepository {
	return &commentCommentRepository{}
}

func (rep *commentCommentRepository) Get(m *model.Comment, db *gorm.DB) ([]model.Comment, error) {
	var accounts []model.Comment
	err := db.Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *commentCommentRepository) GetOne(m *model.Comment, db *gorm.DB) (model.Comment, error) {
	var account model.Comment
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *commentCommentRepository) GetAll(m *model.Comment, db *gorm.DB) ([]model.Comment, error) {
	var accounts []model.Comment
	err := db.Unscoped().Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *commentCommentRepository) Insert(m *model.Comment, db *gorm.DB) (model.Comment, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *commentCommentRepository) Update(m *model.Comment, db *gorm.DB) (model.Comment, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *commentCommentRepository) Delete(m *model.Comment, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}

func (rep *commentCommentRepository) RestoreOne(m *model.Comment, db *gorm.DB) error {
	return db.Unscoped().
		Model(&model.Comment{}).
		Where("id = ?", m.Id).
		Update("deleted_at", nil).Error
}
