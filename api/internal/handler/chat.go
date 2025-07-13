package handler

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/gorilla/websocket"

	"lv99/internal/core"
	"lv99/internal/helper"
	chatModule "lv99/internal/module/chat"
	usecase "lv99/internal/usecase/chat"
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

type PutChatsReadRequest struct {
	FromId int `json:"from_id" binding:"required"`
}

// -----------------------------
// Handler Interface
// -----------------------------

type ChatHandler interface {
	ApiGet(c *gin.Context)
	ApiRead(c *gin.Context)
	WsConnect(c *gin.Context)
}

type chatHandler struct {
	db      *gorm.DB
	usecase usecase.Usecase
}

func NewChatHandler(db *gorm.DB, usecase usecase.Usecase) ChatHandler {
	return &chatHandler{
		db:      db,
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// GET /api/chats/:to_id?before=timestamp
func (ctrl *chatHandler) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri ChatRoomUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	beforeStr := c.Query("before")
	var before time.Time
	if beforeStr == "" {
		before = time.Now()
	} else {
		var err error
		before, err = time.Parse(time.RFC3339Nano, beforeStr)
		if err != nil {
			c.Error(core.ErrBadRequest)
			return
		}
	}

	chats, err := ctrl.usecase.Paginate(usecase.PaginateDto{
		FromId: accountId,
		ToId:   uri.ToId,
		Before: before,
		Limit:  30,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToChatReponseList(chats))
}

// PUT /api/chats/read
func (ctrl *chatHandler) ApiRead(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req PutChatsReadRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.usecase.Read(usecase.ReadDto{
		FromId: req.FromId,
		ToId:   accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}

var sockets = make(map[int][]*websocket.Conn)
var socketsMutex sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 開発用：CORS制限なし
	},
}

func removeConn(accountId int, conn *websocket.Conn) {
	socketsMutex.Lock()
	defer socketsMutex.Unlock()

	conns, ok := sockets[accountId]
	if !ok {
		return
	}

	newConns := make([]*websocket.Conn, 0, len(conns))
	for _, c := range conns {
		if c != conn {
			newConns = append(newConns, c)
		} else {
			_ = c.Close()
		}
	}
	if len(newConns) > 0 {
		sockets[accountId] = newConns
	} else {
		delete(sockets, accountId)
	}
	core.Logger.Info("Removed connection for account %d, remaining: %d", accountId, len(sockets[accountId]))
}

func (ctrl *chatHandler) WsConnect(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		core.Logger.Error("Upgrade error:", err)
		return
	}

	socketsMutex.Lock()
	sockets[accountId] = append(sockets[accountId], conn)
	socketsMutex.Unlock()

	defer func() {
		conn.Close()

		socketsMutex.Lock()
		defer socketsMutex.Unlock()

		conns := sockets[accountId]
		newConns := make([]*websocket.Conn, 0, len(conns))
		for _, c := range conns {
			if c != conn {
				newConns = append(newConns, c)
			}
		}

		if len(newConns) > 0 {
			sockets[accountId] = newConns
		} else {
			delete(sockets, accountId)
		}
		core.Logger.Info("Removed connection for account %d, remaining: %d", accountId, len(sockets[accountId]))
	}()

	for {
		var req GetChatsRequest
		if err := conn.ReadJSON(&req); err != nil {
			core.Logger.Error("ReadJSON error:", err)
			break
		}

		socketsMutex.Lock()
		chat, err := ctrl.usecase.CreateOne(usecase.CreateOneDto{
			FromId:  accountId,
			ToId:    req.ToId,
			Message: req.Message,
		}, ctrl.db)
		if err != nil {
			core.Logger.Error(err.Error())
			socketsMutex.Unlock()
			break
		}

		// toId 宛のコネクションに送信、失敗時に削除
		if toConns, ok := sockets[chat.ToId]; ok {
			for _, toConn := range toConns {
				if err := toConn.WriteJSON(ToChatReponse(chat)); err != nil {
					core.Logger.Warn("Failed to send to %d: %v", chat.ToId, err)
					removeConn(chat.ToId, toConn)
				}
			}
		}

		// fromId 宛のコネクションに送信（toIdと異なる場合）、失敗時に削除
		if chat.FromId != chat.ToId {
			if fromConns, ok := sockets[chat.FromId]; ok {
				for _, fromConn := range fromConns {
					if err := fromConn.WriteJSON(ToChatReponse(chat)); err != nil {
						core.Logger.Warn("Failed to send to %d: %v", chat.FromId, err)
						removeConn(chat.FromId, fromConn)
					}
				}
			}
		}

		socketsMutex.Unlock()
	}
}
