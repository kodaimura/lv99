package chat_extended

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Usecase interface {
	GetUnreadCount(in GetUnreadCountDto) ([]UnreadCount, error)
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

func (uc *usecase) GetUnreadCount(in GetUnreadCountDto) ([]UnreadCount, error) {
	return QueryUnreadCount(in.ToId, uc.dbx)
}
