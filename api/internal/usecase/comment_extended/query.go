package comment_extended

import (
	"time"

	"github.com/jmoiron/sqlx"
)

func QueryWithProfile(answerId int, db *sqlx.DB) ([]CommentWithProfile, error) {
	var accounts []CommentWithProfile

	err := db.Select(&accounts,
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

func QueryCount(accountId int, since *time.Time, db *sqlx.DB) ([]CommentCount, error) {
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
		err = db.Select(&counts, baseQuery, accountId)
	} else {
		err = db.Select(&counts, baseQuery, accountId, since)
	}

	return counts, err
}

func QueryCountForAdmin(accountId int, since *time.Time, db *sqlx.DB) ([]CommentCount, error) {
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
		err = db.Select(&counts, baseQuery, accountId)
	} else {
		err = db.Select(&counts, baseQuery, accountId, since)
	}

	return counts, err
}
