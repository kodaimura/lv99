package answer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Controller interface {
	ApiGetStatus(c *gin.Context)
	AdminGetStatus(c *gin.Context)
	AdminSearch(c *gin.Context)
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

// GET /api/answers/status
func (ctrl *controller) ApiGetStatus(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	status, err := ctrl.service.GetStatus(GetStatusDto{AccountId: accountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerStatusResponseList(status))
}

// GET /api/admin/answers/status?account_id=:account_id
func (ctrl *controller) AdminGetStatus(c *gin.Context) {
	var req GetStatusRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}
	status, err := ctrl.service.GetStatus(GetStatusDto(req))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerStatusResponseList(status))
}

// GET /api/admin/answers/search?...
func (ctrl *controller) AdminSearch(c *gin.Context) {
	var req SearchRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := ctrl.service.Search(SearchDto(req))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerSearchResponseList(answers))
}
