package impl

import (
	"gorm.io/gorm"

	"lv99/internal/model"
)

type gormAnswerRepository struct {
	db *gorm.DB
}

func NewGormAnswerRepository(db *gorm.DB) *gormAnswerRepository {
	return &gormAnswerRepository{db: db}
}

func (rep *gormAnswerRepository) Get(a *model.Answer) ([]model.Answer, error) {
	var answers []model.Answer
	err := rep.db.Order("answer_id ASC").Find(&answers, a).Error
	return answers, handleGormError(err)
}

func (rep *gormAnswerRepository) GetOne(a *model.Answer) (model.Answer, error) {
	var answer model.Answer
	err := rep.db.First(&answer, a).Error
	return answer, handleGormError(err)
}

func (rep *gormAnswerRepository) GetAll(a *model.Answer) ([]model.Answer, error) {
	var answers []model.Answer
	err := rep.db.Unscoped().Order("answer_id ASC").Find(&answers, a).Error
	return answers, handleGormError(err)
}

func (rep *gormAnswerRepository) Insert(a *model.Answer) (model.Answer, error) {
	err := rep.db.Create(a).Error
	return *a, handleGormError(err)
}

func (rep *gormAnswerRepository) Update(a *model.Answer) (model.Answer, error) {
	err := rep.db.Save(a).Error
	return *a, handleGormError(err)
}

func (rep *gormAnswerRepository) Delete(a *model.Answer) error {
	err := rep.db.Delete(a).Error
	return handleGormError(err)
}

func (rep *gormAnswerRepository) InsertTx(a *model.Answer, tx *gorm.DB) (model.Answer, error) {
	err := tx.Create(a).Error
	return *a, handleGormError(err)
}

func (rep *gormAnswerRepository) UpdateTx(a *model.Answer, tx *gorm.DB) (model.Answer, error) {
	err := tx.Save(a).Error
	return *a, handleGormError(err)
}

func (rep *gormAnswerRepository) DeleteTx(a *model.Answer, tx *gorm.DB) error {
	err := tx.Delete(a).Error
	return handleGormError(err)
}
