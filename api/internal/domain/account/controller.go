package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Controller interface {
	ApiGetMe(c *gin.Context)
	ApiPutMe(c *gin.Context)
	ApiDeleteMe(c *gin.Context)
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

// GET /api/accounts/me
func (ctrl *controller) ApiGetMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	account, err := ctrl.service.GetOne(GetOneDto{Id: accountId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountResponse(account))
}

// PUT /api/accounts/me
func (ctrl *controller) ApiPutMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req PutMeRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	account, err := ctrl.service.UpdateOne(UpdateOneDto{
		Id:   accountId,
		Name: req.Name,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountResponse(account))
}

// DELETE /api/accounts/me
func (ctrl *controller) ApiDeleteMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	if err := ctrl.service.DeleteOne(DeleteOneDto{Id: accountId}, ctrl.db); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
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
