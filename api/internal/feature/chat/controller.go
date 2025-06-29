package chat

import (
	"lv99/internal/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller interface {
	ApiGetUnreadCount(c *gin.Context)
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

// GET /api/chats/unread-count
func (ctrl *controller) ApiGetUnreadCount(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	unreadCounts, err := ctrl.service.GetUnreadCount(GetUnreadCountDto{ToId: accountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToUnreadCountResponseList(unreadCounts))
}
