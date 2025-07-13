package comment_extended

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

type CommentCount struct {
	QuestionId    int       `db:"question_id"`
	QuestionTitle string    `db:"question_title"`
	QuestionLevel int       `db:"question_level"`
	AnswerId      int       `db:"answer_id"`
	CommentCount  int       `db:"comment_count"`
	CreatedAt     time.Time `db:"created_at"`
}

type GetCountDto struct {
	AccountId int
	Since     *time.Time
}
