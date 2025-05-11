package query

import (
	"lv99/internal/model"
)

type ChatQuery interface {
	Get(accounId1 int, accountId2 int) ([]model.Chat, error)
}
