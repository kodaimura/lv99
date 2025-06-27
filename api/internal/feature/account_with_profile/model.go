package account_with_profile

import "time"

type AccountWithProfile struct {
	Id          int        `db:"id"`
	Name        string     `db:"name"`
	Role        int        `db:"account_role"`
	DisplayName string     `db:"display_name"`
	Bio         string     `db:"bio"`
	AvatarURL   string     `db:"avatar_url"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
