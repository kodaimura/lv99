package chat_extended

import (
	"github.com/jmoiron/sqlx"
)

func QueryUnreadCount(toId int, db *sqlx.DB) ([]UnreadCount, error) {
	var counts []UnreadCount

	err := db.Select(&counts,
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
