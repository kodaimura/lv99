package chat

import (
	"time"
)

type UnreadCount struct {
	AccountId   int       `db:"account_id"`
	UnreadCount int       `db:"unread_count"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type GetUnreadCountDto struct {
	ToId int
}
