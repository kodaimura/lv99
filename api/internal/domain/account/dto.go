package account

import "time"

type GetDto struct {
	Id   int
	Name string
}

type GetOneDto struct {
	Id int
}

type UpdateOneDto struct {
	Id   int
	Name string
}

type DeleteOneDto struct {
	Id int
}

type GetWithProfileDto struct{}

type AccountWithProfile struct {
	Id          int        `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Role        int        `db:"account_role" json:"role"`
	DisplayName string     `db:"display_name" json:"display_name"`
	Bio         *string    `db:"bio" json:"bio"`
	AvatarURL   *string    `db:"avatar_url" json:"avatar_url"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
}
