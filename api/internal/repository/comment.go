package repository

import (
	"lv99/internal/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Get(c *model.Comment) ([]model.Comment, error)
	GetOne(c *model.Comment) (model.Comment, error)
	GetAll(c *model.Comment) ([]model.Comment, error)

	Insert(c *model.Comment) (model.Comment, error)
	Update(c *model.Comment) (model.Comment, error)
	Delete(c *model.Comment) error

	InsertTx(c *model.Comment, tx *gorm.DB) (model.Comment, error)
	UpdateTx(c *model.Comment, tx *gorm.DB) (model.Comment, error)
	DeleteTx(c *model.Comment, tx *gorm.DB) error
}
