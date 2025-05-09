package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Get(m *model.Comment) ([]model.Comment, error)
	GetOne(m *model.Comment) (model.Comment, error)
	GetAll(m *model.Comment) ([]model.Comment, error)

	Insert(m *model.Comment) (model.Comment, error)
	Update(m *model.Comment) (model.Comment, error)
	Delete(m *model.Comment) error

	InsertTx(m *model.Comment, tx *gorm.DB) (model.Comment, error)
	UpdateTx(m *model.Comment, tx *gorm.DB) (model.Comment, error)
	DeleteTx(m *model.Comment, tx *gorm.DB) error
}
