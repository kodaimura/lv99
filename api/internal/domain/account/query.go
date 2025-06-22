package account

import (
	"github.com/jmoiron/sqlx"
)

type Query interface {
	GetWithProfile() ([]AccountWithProfile, error)
}

type query struct {
	db *sqlx.DB
}

func NewQuery(db *sqlx.DB) *query {
	return &query{db}
}

func (que *query) GetWithProfile() ([]AccountWithProfile, error) {
	var accounts []AccountWithProfile

	err := que.db.Select(&accounts,
		`SELECT
			a.id,
			a.name,
			a.account_role,
			p.display_name,
			p.bio,
			p.avatar_url,
			a.created_at,
			p.updated_at,
			a.deleted_at
		 FROM account as a
		 JOIN account_profile as p 
		   ON a.id = p.account_id
		 ORDER BY a.id DESC`,
	)

	return accounts, err
}
