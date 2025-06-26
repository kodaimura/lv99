package answer_search

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Controller interface {
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
