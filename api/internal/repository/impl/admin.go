package impl

import (
	"gorm.io/gorm"

	"lv99/internal/model"
)

type gormAdminRepository struct {
	db *gorm.DB
}

func NewGormAdminRepository(db *gorm.DB) *gormAdminRepository {
	return &gormAdminRepository{db: db}
}

func (rep *gormAdminRepository) Get(a *model.Admin) ([]model.Admin, error) {
	var admins []model.Admin
	err := rep.db.Find(&admins, a).Error
	return admins, handleGormError(err)
}

func (rep *gormAdminRepository) GetOne(a *model.Admin) (model.Admin, error) {
	var admin model.Admin
	err := rep.db.First(&admin, a).Error
	return admin, handleGormError(err)
}

func (rep *gormAdminRepository) Insert(a *model.Admin) (model.Admin, error) {
	err := rep.db.Create(a).Error
	return *a, handleGormError(err)
}

func (rep *gormAdminRepository) Update(a *model.Admin) (model.Admin, error) {
	err := rep.db.Save(a).Error
	return *a, handleGormError(err)
}

func (rep *gormAdminRepository) Delete(a *model.Admin) error {
	err := rep.db.Delete(a).Error
	return handleGormError(err)
}

func (rep *gormAdminRepository) InsertTx(a *model.Admin, tx *gorm.DB) (model.Admin, error) {
	err := tx.Create(a).Error
	return *a, handleGormError(err)
}

func (rep *gormAdminRepository) UpdateTx(a *model.Admin, tx *gorm.DB) (model.Admin, error) {
	err := tx.Save(a).Error
	return *a, handleGormError(err)
}

func (rep *gormAdminRepository) DeleteTx(a *model.Admin, tx *gorm.DB) error {
	err := tx.Delete(a).Error
	return handleGormError(err)
}
