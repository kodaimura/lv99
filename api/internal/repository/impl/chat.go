package impl

import (
	"gorm.io/gorm"

	"lv99/internal/model"
)

type gormChatRepository struct {
	db *gorm.DB
}

func NewGormChatRepository(db *gorm.DB) *gormChatRepository {
	return &gormChatRepository{db: db}
}

func (rep *gormChatRepository) Get(m *model.Chat) ([]model.Chat, error) {
	var chats []model.Chat
	err := rep.db.Order("id ASC").Find(&chats, m).Error
	return chats, handleGormError(err)
}

func (rep *gormChatRepository) GetOne(m *model.Chat) (model.Chat, error) {
	var chat model.Chat
	err := rep.db.First(&chat, m).Error
	return chat, handleGormError(err)
}

func (rep *gormChatRepository) GetAll(m *model.Chat) ([]model.Chat, error) {
	var chats []model.Chat
	err := rep.db.Unscoped().Order("id ASC").Find(&chats, m).Error
	return chats, handleGormError(err)
}

func (rep *gormChatRepository) Insert(m *model.Chat) (model.Chat, error) {
	err := rep.db.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormChatRepository) Update(m *model.Chat) (model.Chat, error) {
	err := rep.db.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormChatRepository) Delete(m *model.Chat) error {
	err := rep.db.Delete(m).Error
	return handleGormError(err)
}

func (rep *gormChatRepository) InsertTx(m *model.Chat, tx *gorm.DB) (model.Chat, error) {
	err := tx.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormChatRepository) UpdateTx(m *model.Chat, tx *gorm.DB) (model.Chat, error) {
	err := tx.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormChatRepository) DeleteTx(m *model.Chat, tx *gorm.DB) error {
	err := tx.Delete(m).Error
	return handleGormError(err)
}
