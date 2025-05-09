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

func (rep *gormAnswerRepository) Get(m *model.Answer) ([]model.Answer, error) {
	var answers []model.Answer
	err := rep.db.Order("answer_id ASC").Find(&answers, m).Error
	return answers, handleGormError(err)
}

func (rep *gormAnswerRepository) GetOne(m *model.Answer) (model.Answer, error) {
	var answer model.Answer
	err := rep.db.First(&answer, m).Error
	return answer, handleGormError(err)
}

func (rep *gormAnswerRepository) GetAll(m *model.Answer) ([]model.Answer, error) {
	var answers []model.Answer
	err := rep.db.Unscoped().Order("answer_id ASC").Find(&answers, m).Error
	return answers, handleGormError(err)
}

func (rep *gormAnswerRepository) Insert(m *model.Answer) (model.Answer, error) {
	err := rep.db.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormAnswerRepository) Update(m *model.Answer) (model.Answer, error) {
	err := rep.db.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormAnswerRepository) Delete(m *model.Answer) error {
	err := rep.db.Delete(m).Error
	return handleGormError(err)
}

func (rep *gormAnswerRepository) InsertTx(m *model.Answer, tx *gorm.DB) (model.Answer, error) {
	err := tx.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormAnswerRepository) UpdateTx(m *model.Answer, tx *gorm.DB) (model.Answer, error) {
	err := tx.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormAnswerRepository) DeleteTx(m *model.Answer, tx *gorm.DB) error {
	err := tx.Delete(m).Error
	return handleGormError(err)
}
