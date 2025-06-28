package chat

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Query interface {
	Get(accounId1 int, accountId2 int, before time.Time, limit int) ([]Chat, error)
}

type query struct {
	db *sqlx.DB
}

func NewQuery(db *sqlx.DB) *query {
	return &query{db}
}

func (que *query) Get(accounId1 int, accountId2 int, before time.Time, limit int) ([]Chat, error) {
	var chats []Chat

	err := que.db.Select(&chats,
		`SELECT
			id,
			from_id,
			to_id,
			message,
			is_read,
			created_at,
			updated_at
		 FROM chat
		 WHERE ((from_id = $1 AND to_id = $2)
			OR (from_id = $3 AND to_id = $4))
			AND deleted_at IS NULL
			AND created_at < $5
		 ORDER BY created_at DESC
		 LIMIT $6`,
		accounId1,
		accountId2,
		accountId2,
		accounId1,
		before,
		limit,
	)

	return chats, err
}
