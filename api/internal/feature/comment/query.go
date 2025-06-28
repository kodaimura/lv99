package comment

import (
	"github.com/jmoiron/sqlx"
)

type Query interface {
	GetWithProfile(answerId int) ([]CommentWithProfile, error)
}

type query struct {
	db *sqlx.DB
}

func NewQuery(db *sqlx.DB) *query {
	return &query{db}
}

func (que *query) GetWithProfile(answerId int) ([]CommentWithProfile, error) {
	var accounts []CommentWithProfile

	err := que.db.Select(&accounts,
		`SELECT
			c.id,
			c.answer_id,
			c.account_id,
			p.display_name,
			p.avatar_url,
			c.content,
			c.created_at,
			c.updated_at
		 FROM comment as c
		 JOIN account_profile as p 
		   ON c.account_id = p.account_id
		 WHERE c.deleted_at IS NULL
		   AND c.answer_id = $1
		 ORDER BY c.id ASC;`,
		answerId,
	)

	return accounts, err
}
