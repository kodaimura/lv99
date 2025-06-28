package question

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	Id        int            `db:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title     string         `db:"title" gorm:"column:title"`
	Content   string         `db:"content" gorm:"column:content"`
	Answer    string         `db:"answer" gorm:"column:answer"`
	Level     int            `db:"level" gorm:"column:level"`
	CreatedAt time.Time      `db:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"column:deleted_at;index"`
}
