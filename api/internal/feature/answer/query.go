package answer

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Query interface {
	GetStatus(accountId int) ([]AnswerStatus, error)
	Search(accountId, questionId int, isCorrect *bool, commentAccountId int) ([]AnswerSearch, error)
}

type query struct {
	db *sqlx.DB
}

func NewQuery(db *sqlx.DB) *query {
	return &query{db}
}

func (que *query) GetStatus(accountId int) ([]AnswerStatus, error) {
	var answers []AnswerStatus

	baseQuery := `
SELECT
	question_id,
	MAX(CASE WHEN is_correct THEN 1 ELSE 0 END) = 1 AS is_correct,
	COUNT(*) FILTER (WHERE is_correct) AS correct_count,
	MIN(correct_at) FILTER (WHERE is_correct) AS correct_at,
	MAX(updated_at) AS updated_at
FROM answer
WHERE account_id = $1
  AND deleted_at IS NULL
GROUP BY question_id
ORDER BY question_id;`

	err := que.db.Select(&answers, baseQuery, accountId)
	return answers, err
}

func (que *query) Search(accountId, questionId int, isCorrect *bool, commentAccountId int) ([]AnswerSearch, error) {
	var answers []AnswerSearch
	var args []interface{}
	var conditions []string

	baseQuery := `
SELECT
	a.id as answer_id,
	a.code_def as code_def,
	a.code_call as code_call,
	a.is_correct as is_correct,
	a.correct_at as correct_at,
	a.created_at as created_at,
	a.updated_at as updated_at,
	q.id as question_id,
	q.title as question_title,
	q.level as question_level,
	a.account_id as account_id,
	ap.display_name as account_name,
	(SELECT COUNT(*) FROM comment c WHERE c.answer_id = a.id) as comment_count,
	c.account_id as comment_account_id,
	ca.name as comment_account_name,
	c.created_at as comment_at
FROM answer a
JOIN question q ON a.question_id = q.id
JOIN account_profile ap ON a.account_id = ap.account_id
LEFT JOIN LATERAL (
	SELECT * FROM comment
	WHERE comment.answer_id = a.id
	ORDER BY comment.created_at DESC
	LIMIT 1
) c ON true
LEFT JOIN account ca ON ca.id = c.account_id
`
	conditions = append(conditions, "a.deleted_at IS NULL")

	if accountId != 0 {
		conditions = append(conditions, fmt.Sprintf("a.account_id = $%d", len(args)+1))
		args = append(args, accountId)
	}
	if questionId != 0 {
		conditions = append(conditions, fmt.Sprintf("q.id = $%d", len(args)+1))
		args = append(args, questionId)
	}
	if isCorrect != nil {
		conditions = append(conditions, fmt.Sprintf("a.is_correct = $%d", len(args)+1))
		args = append(args, *isCorrect)
	}
	if commentAccountId != 0 {
		conditions = append(conditions, fmt.Sprintf("c.account_id = $%d", len(args)+1))
		args = append(args, commentAccountId)
	}

	if len(conditions) > 0 {
		baseQuery += "WHERE " + strings.Join(conditions, " AND ") + "\n"
	}

	baseQuery += "ORDER BY a.created_at DESC"

	err := que.db.Select(&answers, baseQuery, args...)
	return answers, err
}
