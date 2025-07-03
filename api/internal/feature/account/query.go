package account

import (
	"lv99/internal/helper"

	"github.com/jmoiron/sqlx"
)

type Query interface {
	GetWithProfile() ([]AccountWithProfile, error)
	GetOneWithProfile(accountId int) (AccountWithProfile, error)
	GetAdminWithProfile() (AccountWithProfile, error)
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
		 ORDER BY a.id ASC`,
	)

	return accounts, err
}

func (que *query) GetOneWithProfile(accountId int) (AccountWithProfile, error) {
	var account AccountWithProfile

	err := que.db.Get(&account,
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
		 WHERE a.id = $1`,
		accountId,
	)

	return account, err
}

func (que *query) GetAdminWithProfile() (AccountWithProfile, error) {
	var account AccountWithProfile

	err := que.db.Get(&account,
		`SELECT
			a.id,
			'' as name,
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
		 WHERE a.account_role = $1
		 LIMIT 1`,
		helper.ACCOUNT_ROLE_ADMIN,
	)

	return account, err
}
