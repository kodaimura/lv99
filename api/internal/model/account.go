package model

type Account struct {
	AccountId       int    `db:"account_id" json:"account_id" gorm:"primaryKey;autoIncrement"`
	AccountName     string `db:"account_name" json:"account_name"`
	AccountPassword string `db:"account_password" json:"account_password"`
	AccountRole     int    `db:"account_role" json:"account_role"`
	CreatedAt       string `db:"created_at" json:"created_at" gorm:"-"`
	UpdatedAt       string `db:"updated_at" json:"updated_at" gorm:"-"`
}
