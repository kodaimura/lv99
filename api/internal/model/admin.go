package model

type Admin struct {
	AdminId       int    `db:"admin_id" json:"admin_id" gorm:"primaryKey;autoIncrement"`
	AdminName     string `db:"admin_name" json:"admin_name"`
	AdminPassword string `db:"admin_password" json:"admin_password"`
	CreatedAt     string `db:"created_at" json:"created_at" gorm:"-"`
	UpdatedAt     string `db:"updated_at" json:"updated_at" gorm:"-"`
}
