package comment

import (
	"time"
)

type CommentWithProfile struct {
	Id          int       `db:"id"`
	AnswerId    int       `db:"answer_id"`
	AccountId   int       `db:"account_id"`
	DisplayName string    `db:"display_name"`
	AvatarURL   string    `db:"avatar_url"`
	Content     string    `db:"content"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type GetWithProfileDto struct {
	AnswerId int
}
