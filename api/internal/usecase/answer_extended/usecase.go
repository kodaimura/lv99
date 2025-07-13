package answer_extended

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Usecase interface {
	GetStatus(in GetStatusDto) ([]AnswerStatus, error)
	Search(in SearchDto) ([]AnswerSearch, error)
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

func (uc *usecase) GetStatus(in GetStatusDto) ([]AnswerStatus, error) {
	return QueryStatus(in.AccountId, uc.dbx)
}

func (uc *usecase) Search(in SearchDto) ([]AnswerSearch, error) {
	return QuerySearch(
		in.AccountId,
		in.QuestionId,
		in.Level,
		in.IsCorrect,
		in.CommentAccountId,
		uc.dbx,
	)
}
