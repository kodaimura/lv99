package response

import (
	"lv99/internal/model"
	"time"

	"gorm.io/gorm"
)

// ============================
// Chat
// ============================

type Chat struct {
	Id        int            `json:"id"`
	FromId    int            `json:"from_id"`
	ToId      int            `json:"to_id"`
	Message   string         `json:"message"`
	IsRead    bool           `json:"is_read"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromModelChat(m model.Chat) Chat {
	return Chat{
		Id:        m.Id,
		FromId:    m.FromId,
		ToId:      m.ToId,
		Message:   m.Message,
		IsRead:    m.IsRead,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
	}
}

func FromModelChatList(models []model.Chat) []Chat {
	res := make([]Chat, 0, len(models))
	for _, m := range models {
		res = append(res, FromModelChat(m))
	}
	return res
}
