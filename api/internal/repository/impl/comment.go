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

func (rep *gormCommentRepository) Get(m *model.Comment) ([]model.Comment, error) {
	var comments []model.Comment
	err := rep.db.Order("comment_id ASC").Find(&comments, m).Error
	return comments, handleGormError(err)
}

func (rep *gormCommentRepository) GetOne(m *model.Comment) (model.Comment, error) {
	var comment model.Comment
	err := rep.db.First(&comment, m).Error
	return comment, handleGormError(err)
}

func (rep *gormCommentRepository) GetAll(m *model.Comment) ([]model.Comment, error) {
	var comments []model.Comment
	err := rep.db.Unscoped().Order("comment_id ASC").Find(&comments, m).Error
	return comments, handleGormError(err)
}

func (rep *gormCommentRepository) Insert(m *model.Comment) (model.Comment, error) {
	err := rep.db.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormCommentRepository) Update(m *model.Comment) (model.Comment, error) {
	err := rep.db.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormCommentRepository) Delete(m *model.Comment) error {
	err := rep.db.Delete(m).Error
	return handleGormError(err)
}

func (rep *gormCommentRepository) InsertTx(m *model.Comment, tx *gorm.DB) (model.Comment, error) {
	err := tx.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormCommentRepository) UpdateTx(m *model.Comment, tx *gorm.DB) (model.Comment, error) {
	err := tx.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormCommentRepository) DeleteTx(m *model.Comment, tx *gorm.DB) error {
	err := tx.Delete(m).Error
	return handleGormError(err)
}
