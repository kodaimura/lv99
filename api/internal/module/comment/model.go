package comment

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id        int            `db:"id" gorm:"column:id;primaryKey;autoIncrement"`
	AnswerId  int            `db:"answer_id" gorm:"column:answer_id"`
	AccountId int            `db:"account_id" gorm:"column:account_id"`
	Content   string         `db:"content" gorm:"column:content"`
	CreatedAt time.Time      `db:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `db:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at" gorm:"column:deleted_at;index"`
}
