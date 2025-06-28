package chat

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	Id        int            `db:"id" gorm:"column:id;primaryKey;autoIncrement"`
	FromId    int            `db:"from_id" gorm:"column:from_id"`
	ToId      int            `db:"to_id" gorm:"column:to_id"`
	Message   string         `db:"message" gorm:"column:message"`
	IsRead    bool           `db:"is_read" gorm:"is_read"`
	CreatedAt time.Time      `db:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"column:deleted_at;index"`
}
