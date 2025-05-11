package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type ChatRepository interface {
	Get(m *model.Chat) ([]model.Chat, error)
	GetOne(m *model.Chat) (model.Chat, error)
	GetAll(m *model.Chat) ([]model.Chat, error)

	Insert(m *model.Chat) (model.Chat, error)
	Update(m *model.Chat) (model.Chat, error)
	Delete(m *model.Chat) error

	InsertTx(m *model.Chat, tx *gorm.DB) (model.Chat, error)
	UpdateTx(m *model.Chat, tx *gorm.DB) (model.Chat, error)
	DeleteTx(m *model.Chat, tx *gorm.DB) error
}
