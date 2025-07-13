package chat

import (
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Repository interface {
	Get(m *Chat, db *gorm.DB) ([]Chat, error)
	GetOne(m *Chat, db *gorm.DB) (Chat, error)
	GetAll(m *Chat, db *gorm.DB) ([]Chat, error)
	Insert(m *Chat, db *gorm.DB) (Chat, error)
	Update(m *Chat, db *gorm.DB) (Chat, error)
	Delete(m *Chat, db *gorm.DB) error
	Read(m *Chat, db *gorm.DB) error

	Paginate(accounId1 int, accountId2 int, before time.Time, limit int, db *sqlx.DB) ([]Chat, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (rep *repository) Get(m *Chat, db *gorm.DB) ([]Chat, error) {
	var accounts []Chat
	err := db.Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *repository) GetOne(m *Chat, db *gorm.DB) (Chat, error) {
	var account Chat
	err := db.First(&account, m).Error
	return account, helper.HandleGormError(err)
}

func (rep *repository) GetAll(m *Chat, db *gorm.DB) ([]Chat, error) {
	var accounts []Chat
	err := db.Unscoped().Order("id ASC").Find(&accounts, m).Error
	return accounts, helper.HandleGormError(err)
}

func (rep *repository) Insert(m *Chat, db *gorm.DB) (Chat, error) {
	err := db.Create(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *repository) Update(m *Chat, db *gorm.DB) (Chat, error) {
	err := db.Save(m).Error
	return *m, helper.HandleGormError(err)
}

func (rep *repository) Delete(m *Chat, db *gorm.DB) error {
	err := db.Delete(m).Error
	return helper.HandleGormError(err)
}

func (rep *repository) Read(m *Chat, db *gorm.DB) error {
	err := db.Model(&Chat{}).
		Where(&Chat{ToId: m.ToId, FromId: m.FromId, IsRead: false}).
		Update("is_read", true).Error

	return helper.HandleGormError(err)
}

func (rep *repository) Paginate(accounId1 int, accountId2 int, before time.Time, limit int, db *sqlx.DB) ([]Chat, error) {
	var chats []Chat

	err := db.Select(&chats,
		`SELECT
			id,
			from_id,
			to_id,
			message,
			is_read,
			created_at,
			updated_at
		 FROM chat
		 WHERE ((from_id = $1 AND to_id = $2)
			OR (from_id = $3 AND to_id = $4))
			AND deleted_at IS NULL
			AND created_at < $5
		 ORDER BY created_at DESC
		 LIMIT $6`,
		accounId1,
		accountId2,
		accountId2,
		accounId1,
		before,
		limit,
	)

	return chats, err
}
