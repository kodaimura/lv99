package controller

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"

	"lv99/internal/core"
	"lv99/internal/dto/input"
	"lv99/internal/dto/request"
	"lv99/internal/dto/response"
	"lv99/internal/helper"
	"lv99/internal/service"
)

type ChatController struct {
	chatService service.ChatService
}

func NewChatController(chatService service.ChatService) *ChatController {
	return &ChatController{
		chatService: chatService,
	}
}

// GET /api/chats/:to_id?before=timestamp
func (ctrl *ChatController) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.ChatRoomUri
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

	chats, err := ctrl.chatService.Get(input.GetChat{
		FromId: accountId,
		ToId:   uri.ToId,
		Before: before,
		Limit:  30,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelChatList(chats))
}

/*
// POST /api/answers/:answer_id/chats
func (ctrl *ChatController) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.AnswerUri
	var req request.ChatBody
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	chat, err := ctrl.chatService.CreateOne(input.Chat{
		AnswerId:  uri.AnswerId,
		AccountId: accountId,
		Content:   req.Content,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, response.FromModelChat(chat))
}*/

/*
// GET /api/chats/:chat_id
func (ctrl *ChatController) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.ChatUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	chat, err := ctrl.chatService.GetOne(input.Chat{
		Id:        uri.ChatId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelChat(chat))
}
*/

/*
// PUT /api/chats/:chat_id
func (ctrl *ChatController) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.ChatUri
	var req request.ChatBody
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	chat, err := ctrl.chatService.UpdateOne(input.Chat{
		Id:        uri.ChatId,
		AccountId: accountId,
		Content:   req.Content,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelChat(chat))
}
*/

/*
// DELETE /api/chats/:chat_id
func (ctrl *ChatController) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.ChatUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.chatService.DeleteOne(input.Chat{
		Id:        uri.ChatId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
*/

var sockets = make(map[int]*websocket.Conn)
var socketsMutex sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 開発用：CORS制限なし
	},
}

func (ctrl *ChatController) WsConnect(c *gin.Context) {
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
		var req request.ChatBody
		if err := conn.ReadJSON(&req); err != nil {
			core.Logger.Error("ReadJSON error:", err)
			break
		}

		socketsMutex.Lock()
		chat, err := ctrl.chatService.CreateOne(input.Chat{
			FromId:  accountId,
			ToId:    req.ToId,
			Message: req.Message,
		})
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
			_ = toConn.WriteJSON(response.FromModelChat(chat))
		}
		if chat.FromId != chat.ToId {
			if fromConn, ok := sockets[chat.FromId]; ok {
				_ = fromConn.WriteJSON(response.FromModelChat(chat))
			}
		}
		socketsMutex.Unlock()
	}
}
