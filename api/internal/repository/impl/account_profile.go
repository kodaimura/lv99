package impl

import (
	"gorm.io/gorm"

	"lv99/internal/model"
)

type gormAccountProfileRepository struct {
	db *gorm.DB
}

func NewGormAccountProfileRepository(db *gorm.DB) *gormAccountProfileRepository {
	return &gormAccountProfileRepository{db: db}
}

func (rep *gormAccountProfileRepository) Get(m *model.AccountProfile) ([]model.AccountProfile, error) {
	var profiles []model.AccountProfile
	err := rep.db.Find(&profiles, m).Error
	return profiles, handleGormError(err)
}

func (rep *gormAccountProfileRepository) GetOne(m *model.AccountProfile) (model.AccountProfile, error) {
	var profile model.AccountProfile
	err := rep.db.First(&profile, m).Error
	return profile, handleGormError(err)
}

func (rep *gormAccountProfileRepository) GetAll(m *model.AccountProfile) ([]model.AccountProfile, error) {
	var profiles []model.AccountProfile
	err := rep.db.Unscoped().Find(&profiles, m).Error
	return profiles, handleGormError(err)
}

func (rep *gormAccountProfileRepository) Insert(m *model.AccountProfile) (model.AccountProfile, error) {
	err := rep.db.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormAccountProfileRepository) Update(m *model.AccountProfile) (model.AccountProfile, error) {
	err := rep.db.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormAccountProfileRepository) Delete(m *model.AccountProfile) error {
	err := rep.db.Delete(m).Error
	return handleGormError(err)
}

func (rep *gormAccountProfileRepository) InsertTx(m *model.AccountProfile, tx *gorm.DB) (model.AccountProfile, error) {
	err := tx.Create(m).Error
	return *m, handleGormError(err)
}

func (rep *gormAccountProfileRepository) UpdateTx(m *model.AccountProfile, tx *gorm.DB) (model.AccountProfile, error) {
	err := tx.Save(m).Error
	return *m, handleGormError(err)
}

func (rep *gormAccountProfileRepository) DeleteTx(m *model.AccountProfile, tx *gorm.DB) error {
	err := tx.Delete(m).Error
	return handleGormError(err)
}
