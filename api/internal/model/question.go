package model

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	Id        int            `db:"id" json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string         `db:"title" json:"title"`
	Content   string         `db:"content" json:"content"`
	Answer    string         `db:"answer" json:"answer"`
	Level     int            `db:"level" json:"level"`
	CreatedAt time.Time      `db:"created_at" json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" json:"deleted_at" gorm:"index"`
}
