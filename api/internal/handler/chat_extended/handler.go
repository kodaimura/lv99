package chat_extended

import (
	"lv99/internal/helper"
	usecase "lv99/internal/usecase/chat_extended"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	ApiGetUnreadCount(c *gin.Context)
}

type handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) Handler {
	return &handler{
		usecase: usecase,
	}
}

// GET /api/chats/unread-count
func (h *handler) ApiGetUnreadCount(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	unreadCounts, err := h.usecase.GetUnreadCount(usecase.GetUnreadCountDto{ToId: accountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToUnreadCountResponseList(unreadCounts))
}
