package account_extended

import (
	"lv99/internal/helper"

	"github.com/jmoiron/sqlx"
)

func QueryWithProfile(db *sqlx.DB) ([]AccountWithProfile, error) {
	var accounts []AccountWithProfile

	err := db.Select(&accounts,
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

func QueryOneWithProfile(accountId int, db *sqlx.DB) (AccountWithProfile, error) {
	var account AccountWithProfile

	err := db.Get(&account,
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

func QueryAdminWithProfile(db *sqlx.DB) (AccountWithProfile, error) {
	var account AccountWithProfile

	err := db.Get(&account,
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
