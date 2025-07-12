package repository

import (
	"gorm.io/gorm"

	"lv99/internal/helper"
	"lv99/internal/model"
)

type ChatRepository interface {
	Get(m *model.Chat, db *gorm.DB) ([]model.Chat, error)
	GetOne(m *model.Chat, db *gorm.DB) (model.Chat, error)
	GetAll(m *model.Chat, db *gorm.DB) ([]model.Chat, error)
	Insert(m *model.Chat, db *gorm.DB) (model.Chat, error)
	Update(m *model.Chat, db *gorm.DB) (model.Chat, error)
	Delete(m *model.Chat, db *gorm.DB) error
	Read(m *model.Chat, db *gorm.DB) error
}

type chatRepository struct{}

func NewChatRepository() ChatRepository {
	return &chatRepository{}
}

func (rep *chatRepository) Get(m *model.Chat, db *gorm.DB) ([]model.Chat, error) {
	var accounts []model.Chat
	err := db.Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *chatRepository) GetOne(m *model.Chat, db *gorm.DB) (model.Chat, error) {
	var account model.Chat
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *chatRepository) GetAll(m *model.Chat, db *gorm.DB) ([]model.Chat, error) {
	var accounts []model.Chat
	err := db.Unscoped().Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *chatRepository) Insert(m *model.Chat, db *gorm.DB) (model.Chat, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *chatRepository) Update(m *model.Chat, db *gorm.DB) (model.Chat, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *chatRepository) Delete(m *model.Chat, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}

func (rep *chatRepository) Read(m *model.Chat, db *gorm.DB) error {
	err := db.Model(&model.Chat{}).
		Where(&model.Chat{ToId: m.ToId, FromId: m.FromId, IsRead: false}).
		Update("is_read", true).Error

	return helper.HandleGormError(err)
}
