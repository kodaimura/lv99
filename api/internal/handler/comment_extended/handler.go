package comment_extended

import (
	"lv99/internal/helper"
	usecase "lv99/internal/usecase/comment_extended"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	ApiGetWithProfile(c *gin.Context)
	ApiGetRecentCount(c *gin.Context)
	AdminGetRecentCount(c *gin.Context)
}

type controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) Controller {
	return &controller{
		usecase: usecase,
	}
}

// GET /api/api/comments/with-profile?answer_id=:answer_id
func (ctrl *controller) ApiGetWithProfile(c *gin.Context) {
	var req GetWithProfileRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}
	comments, err := ctrl.usecase.GetWithProfile(usecase.GetWithProfileDto(req))
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

	counts, err := ctrl.usecase.GetCount(usecase.GetCountDto{
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

	counts, err := ctrl.usecase.GetCountForAdmin(usecase.GetCountDto{
		AccountId: accountId,
		Since:     req.Since,
	})
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, ToCommentCountResponseList(counts))
}
