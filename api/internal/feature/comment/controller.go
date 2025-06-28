package comment

import (
	"lv99/internal/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller interface {
	AdminGetWithProfile(c *gin.Context)
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

// GET /api/admin/comments/with-profile?answer_id=:answer_id
func (ctrl *controller) AdminGetWithProfile(c *gin.Context) {
	var req GetWithProfileRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}
	comments, err := ctrl.service.GetWithProfile(GetWithProfileDto(req))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToCommentWithProfileResponseList(comments))
}
