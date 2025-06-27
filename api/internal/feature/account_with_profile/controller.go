package account_with_profile

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Controller interface {
	AdminGet(c *gin.Context)
	AdminGetOne(c *gin.Context)
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

// GET /api/admin/accounts/with-profile
func (ctrl *controller) AdminGet(c *gin.Context) {
	accounts, err := ctrl.service.Get(GetDto{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountWithProfileResponseList(accounts))
}

// GET /api/admin/accounts/:account_id/with-profile
func (ctrl *controller) AdminGetOne(c *gin.Context) {
	var uri AccountUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	account, err := ctrl.service.GetOne(GetOneDto{Id: uri.AccountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountWithProfileResponse(account))
}
