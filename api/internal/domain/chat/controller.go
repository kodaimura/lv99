package chat

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/gorilla/websocket"

	"lv99/internal/core"
	"lv99/internal/helper"
)

type Controller interface {
	ApiGet(c *gin.Context)
	WsConnect(c *gin.Context)
}

type controller struct {
	db      *gorm.DB
	service Service
}

func NewController(db *gorm.DB, service Service) Controller {
	return &controller{
		db:      db,
		service: service,
	}
}

// GET /api/chats/:to_id?before=timestamp
func (ctrl *controller) ApiGet(c *gin.Context) {
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

	chats, err := ctrl.service.Get(GetDto{
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

var sockets = make(map[int]*websocket.Conn)
var socketsMutex sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 開発用：CORS制限なし
	},
}

func (ctrl *controller) WsConnect(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		core.Logger.Error("Upgrade error:", err)
		return
	}

	socketsMutex.Lock()
	sockets[accountId] = conn
	socketsMutex.Unlock()

	defer func() {
		socketsMutex.Lock()
		delete(sockets, accountId)
		socketsMutex.Unlock()
		conn.Close()
	}()

	for {
		var req ChatRequest
		if err := conn.ReadJSON(&req); err != nil {
			core.Logger.Error("ReadJSON error:", err)
			break
		}

		socketsMutex.Lock()
		chat, err := ctrl.service.CreateOne(CreateOneDto{
			FromId:  accountId,
			ToId:    req.ToId,
			Message: req.Message,
		}, ctrl.db)
		if err != nil {
			core.Logger.Error(err.Error())
			if toConn, ok := sockets[chat.FromId]; ok {
				_ = toConn.WriteJSON(gin.H{
					"error": "送信に失敗しました。",
				})
			}
			break
		}

		if toConn, ok := sockets[chat.ToId]; ok {
			_ = toConn.WriteJSON(ToChatReponse(chat))
		}
		if chat.FromId != chat.ToId {
			if fromConn, ok := sockets[chat.FromId]; ok {
				_ = fromConn.WriteJSON(ToChatReponse(chat))
			}
		}
		socketsMutex.Unlock()
	}
}
