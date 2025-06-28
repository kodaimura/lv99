package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Controller interface {
	AdminGetWithProfile(c *gin.Context)
	AdminGetOneWithProfile(c *gin.Context)
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
func (ctrl *controller) AdminGetWithProfile(c *gin.Context) {
	accounts, err := ctrl.service.GetWithProfile(GetWithProfileDto{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountWithProfileResponseList(accounts))
}

// GET /api/admin/accounts/:account_id/with-profile
func (ctrl *controller) AdminGetOneWithProfile(c *gin.Context) {
	var uri AccountUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	account, err := ctrl.service.GetOneWithProfile(GetOneWithProfileDto{Id: uri.AccountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountWithProfileResponse(account))
}
