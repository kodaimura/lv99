package query

import (
	"lv99/internal/model"
	"time"
)

type ChatQuery interface {
	Get(accounId1 int, accountId2 int, before time.Time, limit int) ([]model.Chat, error)
}
