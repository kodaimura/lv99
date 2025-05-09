package impl

import (
	"gorm.io/gorm"

	"lv99/internal/model"
)

type gormCommentRepository struct {
	db *gorm.DB
}

func NewGormCommentRepository(db *gorm.DB) *gormCommentRepository {
	return &gormCommentRepository{db: db}
}

func (rep *gormCommentRepository) Get(c *model.Comment) ([]model.Comment, error) {
	var comments []model.Comment
	err := rep.db.Order("comment_id ASC").Find(&comments, c).Error
	return comments, handleGormError(err)
}

func (rep *gormCommentRepository) GetOne(c *model.Comment) (model.Comment, error) {
	var comment model.Comment
	err := rep.db.First(&comment, c).Error
	return comment, handleGormError(err)
}

func (rep *gormCommentRepository) GetAll(c *model.Comment) ([]model.Comment, error) {
	var comments []model.Comment
	err := rep.db.Unscoped().Order("comment_id ASC").Find(&comments, c).Error
	return comments, handleGormError(err)
}

func (rep *gormCommentRepository) Insert(c *model.Comment) (model.Comment, error) {
	err := rep.db.Create(c).Error
	return *c, handleGormError(err)
}

func (rep *gormCommentRepository) Update(c *model.Comment) (model.Comment, error) {
	err := rep.db.Save(c).Error
	return *c, handleGormError(err)
}

func (rep *gormCommentRepository) Delete(c *model.Comment) error {
	err := rep.db.Delete(c).Error
	return handleGormError(err)
}

func (rep *gormCommentRepository) InsertTx(c *model.Comment, tx *gorm.DB) (model.Comment, error) {
	err := tx.Create(c).Error
	return *c, handleGormError(err)
}

func (rep *gormCommentRepository) UpdateTx(c *model.Comment, tx *gorm.DB) (model.Comment, error) {
	err := tx.Save(c).Error
	return *c, handleGormError(err)
}

func (rep *gormCommentRepository) DeleteTx(c *model.Comment, tx *gorm.DB) error {
	err := tx.Delete(c).Error
	return handleGormError(err)
}
