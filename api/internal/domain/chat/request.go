package chat

type ChatUri struct {
	ChatId int `uri:"chat_id" binding:"required"`
}

type ChatRoomUri struct {
	ToId int `uri:"to_id" binding:"required"`
}

type ChatRequest struct {
	ToId    int    `json:"to_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}
