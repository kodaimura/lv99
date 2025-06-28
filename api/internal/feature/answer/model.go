package answer

import (
	"time"
)

type AnswerSearch struct {
	AnswerId           int        `db:"answer_id"`
	CodeDef            string     `db:"code_def"`
	CodeCall           string     `db:"code_call"`
	IsCorrect          bool       `db:"is_correct"`
	CorrectAt          *time.Time `db:"correct_at"`
	CreatedAt          time.Time  `db:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at"`
	QuestionId         int        `db:"question_id"`
	QuestionTitle      string     `db:"question_title"`
	QuestionLevel      int        `db:"question_level"`
	AccountId          int        `db:"account_id"`
	AccountName        string     `db:"account_name"`
	CommentCount       int        `db:"comment_count"`
	CommentAccountId   *int       `db:"comment_account_id"`
	CommentAccountName *string    `db:"comment_account_name"`
	CommentAt          *time.Time `db:"comment_at"`
}
