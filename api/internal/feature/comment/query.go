package comment

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Query interface {
	GetWithProfile(answerId int) ([]CommentWithProfile, error)
	GetCount(accountId int, since *time.Time) ([]CommentCount, error)
	GetCountForAdmin(accountId int, since *time.Time) ([]CommentCount, error)
}

type query struct {
	db *sqlx.DB
}

func NewQuery(db *sqlx.DB) *query {
	return &query{db}
}

func (que *query) GetWithProfile(answerId int) ([]CommentWithProfile, error) {
	var accounts []CommentWithProfile

	err := que.db.Select(&accounts,
		`SELECT
			c.id,
			c.answer_id,
			c.account_id,
			p.display_name,
			p.avatar_url,
			c.content,
			c.created_at,
			c.updated_at
		 FROM comment as c
		 JOIN account_profile as p 
		   ON c.account_id = p.account_id
		 WHERE c.deleted_at IS NULL
		   AND c.answer_id = $1
		 ORDER BY c.id ASC;`,
		answerId,
	)

	return accounts, err
}

func (que *query) GetCount(accountId int, since *time.Time) ([]CommentCount, error) {
	var counts []CommentCount

	baseQuery := `
		SELECT
			q.id AS question_id,
			q.title AS question_title,
			q.level AS question_level,
			a.id AS answer_id,
			COUNT(c.id) AS comment_count,
			COALESCE(MAX(c.created_at), a.created_at) AS created_at
		FROM answer a
		JOIN question q ON q.id = a.question_id
		LEFT JOIN comment c ON c.answer_id = a.id
			AND c.account_id != $1
			AND c.deleted_at IS NULL
		WHERE a.account_id = $1
		  AND a.deleted_at IS NULL
		  AND q.deleted_at IS NULL
	`

	if since != nil {
		baseQuery += ` AND c.created_at >= $2`
	}

	baseQuery += `
		GROUP BY q.id, q.title, q.level, a.id
		HAVING COUNT(c.id) >= 1
		ORDER BY MAX(c.created_at) DESC;
	`

	var err error
	if since == nil {
		err = que.db.Select(&counts, baseQuery, accountId)
	} else {
		err = que.db.Select(&counts, baseQuery, accountId, since)
	}

	return counts, err
}

func (que *query) GetCountForAdmin(accountId int, since *time.Time) ([]CommentCount, error) {
	var counts []CommentCount

	baseQuery := `
		SELECT
			q.id AS question_id,
			q.title AS question_title,
			q.level AS question_level,
			a.id AS answer_id,
			COUNT(c.id) AS comment_count,
			COALESCE(MAX(c.created_at), a.created_at) AS created_at
		FROM answer a
		JOIN question q ON q.id = a.question_id
		LEFT JOIN comment c ON c.answer_id = a.id
			AND c.account_id != $1
			AND c.deleted_at IS NULL
		WHERE a.deleted_at IS NULL
		  AND q.deleted_at IS NULL
	`

	if since != nil {
		baseQuery += ` AND c.created_at >= $2`
	}

	baseQuery += `
		GROUP BY q.id, q.title, q.level, a.id
		HAVING COUNT(c.id) >= 1
		ORDER BY MAX(c.created_at) DESC;
	`

	var err error
	if since == nil {
		err = que.db.Select(&counts, baseQuery, accountId)
	} else {
		err = que.db.Select(&counts, baseQuery, accountId, since)
	}

	return counts, err
}
