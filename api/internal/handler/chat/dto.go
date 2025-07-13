package chat

import (
	"time"

	"gorm.io/gorm"

	chatModule "lv99/internal/module/chat"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type ChatResponse struct {
	Id        int            `json:"id"`
	FromId    int            `json:"from_id"`
	ToId      int            `json:"to_id"`
	Message   string         `json:"message"`
	IsRead    bool           `json:"is_read"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func ToChatReponse(m chatModule.Chat) ChatResponse {
	return ChatResponse{
		Id:        m.Id,
		FromId:    m.FromId,
		ToId:      m.ToId,
		Message:   m.Message,
		IsRead:    m.IsRead,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToChatReponseList(models []chatModule.Chat) []ChatResponse {
	res := make([]ChatResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToChatReponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type ChatUri struct {
	ChatId int `uri:"chat_id" binding:"required"`
}

type ChatRoomUri struct {
	ToId int `uri:"to_id" binding:"required"`
}

type GetChatsRequest struct {
	ToId    int    `json:"to_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type PutReadRequest struct {
	FromId int `json:"from_id" binding:"required"`
}
