package chat

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"

	"lv99/internal/core"
	"lv99/internal/helper"
	usecase "lv99/internal/usecase/chat"
)

// -----------------------------
// Handler Interface
// -----------------------------

type Handler interface {
	ApiGet(c *gin.Context)
	ApiRead(c *gin.Context)
	WsConnect(c *gin.Context)
}

type handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) Handler {
	return &handler{
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// GET /api/chats/:to_id?before=timestamp
func (h *handler) ApiGet(c *gin.Context) {
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

	chats, err := h.usecase.Paginate(usecase.PaginateDto{
		FromId: accountId,
		ToId:   uri.ToId,
		Before: before,
		Limit:  30,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToChatReponseList(chats))
}

// PUT /api/chats/read
func (h *handler) ApiRead(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req PutReadRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := h.usecase.Read(usecase.ReadDto{
		FromId: req.FromId,
		ToId:   accountId,
	})
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

func (h *handler) WsConnect(c *gin.Context) {
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
		chat, err := h.usecase.CreateOne(usecase.CreateOneDto{
			FromId:  accountId,
			ToId:    req.ToId,
			Message: req.Message,
		})
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
