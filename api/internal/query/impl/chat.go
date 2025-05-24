package query

import (
	"lv99/internal/model"

	"github.com/jmoiron/sqlx"
)

type chatQuery struct {
	db *sqlx.DB
}

func NewChatQuery(db *sqlx.DB) *chatQuery {
	return &chatQuery{db}
}

func (que *chatQuery) Get(accounId1 int, accountId2 int) ([]model.Chat, error) {
	var chats []model.Chat

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
		 WHERE (from_id = $1 AND to_id = $2)
			OR (from_id = $3 AND to_id = $4)
		 ORDER BY created_at`,
		accounId1,
		accountId2,
		accountId2,
		accounId1,
	)

	return chats, err
}
