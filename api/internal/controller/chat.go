package controller

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"

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

// GET /api/chats
func (ctrl *ChatController) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	chats, err := ctrl.chatService.Get(input.Chat{
		FromId: accountId,
		ToId: 1,
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
		log.Println("Upgrade error:", err)
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
			log.Println("ReadJSON error:", err)
			break
		}
		log.Printf("From %d to %d: %s", accountId, req.ToId, req.Message)

		socketsMutex.Lock()
		if toConn, ok := sockets[req.ToId]; ok {
			_ = toConn.WriteJSON(gin.H{
				"from_id": accountId,
				"message": req.Message,
			})
		}
		socketsMutex.Unlock()
	}
}