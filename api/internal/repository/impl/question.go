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

func (rep *gormQuestionRepository) Get(q *model.Question) ([]model.Question, error) {
	var questions []model.Question
	err := rep.db.Find(&questions, q).Error
	return questions, handleGormError(err)
}

func (rep *gormQuestionRepository) GetOne(q *model.Question) (model.Question, error) {
	var question model.Question
	err := rep.db.First(&question, q).Error
	return question, handleGormError(err)
}

func (rep *gormQuestionRepository) GetAll(q *model.Question) ([]model.Question, error) {
	var questions []model.Question
	err := rep.db.Unscoped().Find(&questions, q).Error
	return questions, handleGormError(err)
}

func (rep *gormQuestionRepository) Insert(q *model.Question) (model.Question, error) {
	err := rep.db.Create(q).Error
	return *q, handleGormError(err)
}

func (rep *gormQuestionRepository) Update(q *model.Question) (model.Question, error) {
	err := rep.db.Save(q).Error
	return *q, handleGormError(err)
}

func (rep *gormQuestionRepository) Delete(q *model.Question) error {
	err := rep.db.Delete(q).Error
	return handleGormError(err)
}

func (rep *gormQuestionRepository) InsertTx(q *model.Question, tx *gorm.DB) (model.Question, error) {
	err := tx.Create(q).Error
	return *q, handleGormError(err)
}

func (rep *gormQuestionRepository) UpdateTx(q *model.Question, tx *gorm.DB) (model.Question, error) {
	err := tx.Save(q).Error
	return *q, handleGormError(err)
}

func (rep *gormQuestionRepository) DeleteTx(q *model.Question, tx *gorm.DB) error {
	err := tx.Delete(q).Error
	return handleGormError(err)
}

func (rep *gormQuestionRepository) RestoreOne(q *model.Question) error {
	return rep.db.
		Unscoped().
		Model(&model.Question{}).
		Where("question_id = ?", q.QuestionId).
		Update("deleted_at", nil).Error
}