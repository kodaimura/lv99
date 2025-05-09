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

func (rep *gormQuestionRepository) Get(m *model.Question) ([]model.Question, error) {
	var questions []model.Question
	err := rep.db.Order("question_id ASC").Find(&questions, m).Error
	return questions, handleGormError(err)
}

func (rep *gormQuestionRepository) GetOne(m *model.Question) (model.Question, error) {
	var question model.Question
	err := rep.db.First(&question, m).Error
	return question, handleGormError(err)
}

func (rep *gormQuestionRepository) GetAll(m *model.Question) ([]model.Question, error) {
	var questions []model.Question
	err := rep.db.Unscoped().Order("question_id ASC").Find(&questions, m).Error
	return questions, handleGormError(err)
}

func (rep *gormQuestionRepository) Insert(m *model.Question) (model.Question, error) {
	err := rep.db.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormQuestionRepository) Update(m *model.Question) (model.Question, error) {
	err := rep.db.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormQuestionRepository) Delete(m *model.Question) error {
	err := rep.db.Delete(m).Error
	return handleGormError(err)
}

func (rep *gormQuestionRepository) InsertTx(m *model.Question, tx *gorm.DB) (model.Question, error) {
	err := tx.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormQuestionRepository) UpdateTx(m *model.Question, tx *gorm.DB) (model.Question, error) {
	err := tx.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormQuestionRepository) DeleteTx(m *model.Question, tx *gorm.DB) error {
	err := tx.Delete(m).Error
	return handleGormError(err)
}

func (rep *gormQuestionRepository) RestoreOne(m *model.Question) error {
	return rep.db.
		Unscoped().
		Model(&model.Question{}).
		Where("question_id = ?", m.QuestionId).
		Update("deleted_at", nil).Error
}