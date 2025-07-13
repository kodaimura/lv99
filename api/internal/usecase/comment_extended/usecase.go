package comment_extended

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Usecase interface {
	GetWithProfile(in GetWithProfileDto) ([]CommentWithProfile, error)
	GetCount(in GetCountDto) ([]CommentCount, error)
	GetCountForAdmin(in GetCountDto) ([]CommentCount, error)
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

func (uc *usecase) GetWithProfile(in GetWithProfileDto) ([]CommentWithProfile, error) {
	return QueryWithProfile(in.AnswerId, uc.dbx)
}

func (uc *usecase) GetCount(in GetCountDto) ([]CommentCount, error) {
	return QueryCount(in.AccountId, in.Since, uc.dbx)
}

func (uc *usecase) GetCountForAdmin(in GetCountDto) ([]CommentCount, error) {
	return QueryCountForAdmin(in.AccountId, in.Since, uc.dbx)
}
