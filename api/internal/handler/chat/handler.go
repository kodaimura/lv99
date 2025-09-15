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

const (
	writeWait  = 10 * time.Second
	pongWait   = 30 * time.Second
	pingPeriod = 27 * time.Second
)

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

	conn.SetCloseHandler(func(code int, text string) error {
		removeConn(accountId, conn)
		return nil
	})

	// pong で ReadDeadline を延長
	_ = conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(appData string) error {
		_ = conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// 定期 ping（書き込みはロック下で行う）
	stopPing := make(chan struct{})
	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				socketsMutex.Lock()
				_ = conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeWait)); err != nil {
					socketsMutex.Unlock()
					removeConn(accountId, conn)
					return
				}
				socketsMutex.Unlock()
			case <-stopPing:
				return
			}
		}
	}()

	socketsMutex.Lock()
	sockets[accountId] = append(sockets[accountId], conn)
	socketsMutex.Unlock()

	defer func() {
		close(stopPing)
		removeConn(accountId, conn)
	}()

	for {
		var req GetChatsRequest
		if err := conn.ReadJSON(&req); err != nil {
			core.Logger.Error("ReadJSON error:", err)
			break
		}

		chat, err := h.usecase.CreateOne(usecase.CreateOneDto{
			FromId:  accountId,
			ToId:    req.ToId,
			Message: req.Message,
		})
		if err != nil {
			core.Logger.Error(err.Error())
			break
		}

		// 宛先/送信者の接続リストをスナップショット
		var toConns []*websocket.Conn
		var fromConns []*websocket.Conn

		socketsMutex.Lock()
		if list, ok := sockets[chat.ToId]; ok {
			toConns = append(toConns, list...)
		}
		if chat.FromId != chat.ToId {
			if list, ok := sockets[chat.FromId]; ok {
				fromConns = append(fromConns, list...)
			}
		}

		resp := ToChatReponse(chat)

		// ロック下で送信＆失敗ならその場で map から除去（同時書き込み防止）
		for _, toConn := range toConns {
			_ = toConn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := toConn.WriteJSON(resp); err != nil {
				core.Logger.Warn("Failed to send to %d: %v", chat.ToId, err)
				cur := sockets[chat.ToId]
				kept := cur[:0]
				for _, c := range cur {
					if c != toConn {
						kept = append(kept, c)
					} else {
						_ = c.Close()
					}
				}
				if len(kept) > 0 {
					sockets[chat.ToId] = kept
				} else {
					delete(sockets, chat.ToId)
				}
			}
		}

		if len(fromConns) > 0 {
			for _, fromConn := range fromConns {
				_ = fromConn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := fromConn.WriteJSON(resp); err != nil {
					core.Logger.Warn("Failed to send to %d: %v", chat.FromId, err)
					cur := sockets[chat.FromId]
					kept := cur[:0]
					for _, c := range cur {
						if c != fromConn {
							kept = append(kept, c)
						} else {
							_ = c.Close()
						}
					}
					if len(kept) > 0 {
						sockets[chat.FromId] = kept
					} else {
						delete(sockets, chat.FromId)
					}
				}
			}
		}
		socketsMutex.Unlock()
	}
}
