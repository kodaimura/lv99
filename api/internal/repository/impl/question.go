package impl

import (
	"gorm.io/gorm"

	"lv99/internal/model"
)

type gormQuestionRepository struct {
	db *gorm.DB
}

func NewGormQuestionRepository(db *gorm.DB) *gormQuestionRepository {
	return &gormQuestionRepository{db: db}
}

func (rep *gormQuestionRepository) Get(a *model.Question) ([]model.Question, error) {
	var accounts []model.Question
	err := rep.db.Find(&accounts, a).Error
	return accounts, handleGormError(err)
}

func (rep *gormQuestionRepository) GetOne(a *model.Question) (model.Question, error) {
	var account model.Question
	err := rep.db.First(&account, a).Error
	return account, handleGormError(err)
}

func (rep *gormQuestionRepository) Insert(a *model.Question) (model.Question, error) {
	err := rep.db.Create(a).Error
	return *a, handleGormError(err)
}

func (rep *gormQuestionRepository) Update(a *model.Question) (model.Question, error) {
	err := rep.db.Save(a).Error
	return *a, handleGormError(err)
}

func (rep *gormQuestionRepository) Delete(a *model.Question) error {
	err := rep.db.Delete(a).Error
	return handleGormError(err)
}

func (rep *gormQuestionRepository) InsertTx(a *model.Question, tx *gorm.DB) (model.Question, error) {
	err := tx.Create(a).Error
	return *a, handleGormError(err)
}

func (rep *gormQuestionRepository) UpdateTx(a *model.Question, tx *gorm.DB) (model.Question, error) {
	err := tx.Save(a).Error
	return *a, handleGormError(err)
}

func (rep *gormQuestionRepository) DeleteTx(a *model.Question, tx *gorm.DB) error {
	err := tx.Delete(a).Error
	return handleGormError(err)
}
