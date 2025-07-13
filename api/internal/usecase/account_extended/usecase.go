package account_extended

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Usecase interface {
	GetWithProfile(in GetWithProfileDto) ([]AccountWithProfile, error)
	GetOneWithProfile(in GetOneWithProfileDto) (AccountWithProfile, error)
	GetAdminWithProfile() (AccountWithProfile, error)
}

type usecase struct {
	db  *gorm.DB
	dbx *sqlx.DB
}

func NewUsecase(db *gorm.DB, dbx *sqlx.DB) Usecase {
	return &usecase{
		db:  db,
		dbx: dbx,
	}
}

func (uc *usecase) GetWithProfile(in GetWithProfileDto) ([]AccountWithProfile, error) {
	return QueryWithProfile(uc.dbx)
}

func (uc *usecase) GetOneWithProfile(in GetOneWithProfileDto) (AccountWithProfile, error) {
	return QueryOneWithProfile(in.Id, uc.dbx)
}

func (uc *usecase) GetAdminWithProfile() (AccountWithProfile, error) {
	return QueryAdminWithProfile(uc.dbx)
}
