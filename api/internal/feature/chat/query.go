package chat

import (
	"github.com/jmoiron/sqlx"
)

type Query interface {
	GetUnreadCount(toId int) ([]UnreadCount, error)
}

type query struct {
	db *sqlx.DB
}

func NewQuery(db *sqlx.DB) *query {
	return &query{db}
}

func (que *query) GetUnreadCount(toId int) ([]UnreadCount, error) {
	var counts []UnreadCount

	err := que.db.Select(&counts,
		`SELECT
			from_id AS account_id,
			MAX(updated_at) AS updated_at,
			COUNT(*) AS unread_count
		 FROM chat
		 WHERE to_id = $1
		   AND is_read = false
		   AND deleted_at IS NULL
		 GROUP BY from_id
		 ORDER BY updated_at DESC;`,
		toId,
	)

	return counts, err
}
