package comment

import (
	"lv99/internal/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller interface {
	ApiGetWithProfile(c *gin.Context)
	ApiGetRecentCount(c *gin.Context)
	AdminGetRecentCount(c *gin.Context)
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

// GET /api/api/comments/with-profile?answer_id=:answer_id
func (ctrl *controller) ApiGetWithProfile(c *gin.Context) {
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

// GET /api/comments/count?since=:since
func (ctrl *controller) ApiGetRecentCount(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req GetCountRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	counts, err := ctrl.service.GetCount(GetCountDto{
		AccountId: accountId,
		Since:     req.Since,
	})
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, ToCommentCountResponseList(counts))
}

// GET /api/admin/comments/count?since=:since
func (ctrl *controller) AdminGetRecentCount(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req GetCountRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	counts, err := ctrl.service.GetCountForAdmin(GetCountDto{
		AccountId: accountId,
		Since:     req.Since,
	})
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, ToCommentCountResponseList(counts))
}
