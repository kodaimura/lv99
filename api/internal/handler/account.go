package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
	module "lv99/internal/module/account"
)

type AccountHandler interface {
	ApiGetMe(c *gin.Context)
	ApiPutMe(c *gin.Context)
	ApiDeleteMe(c *gin.Context)
}

type accountHandler struct {
	db      *gorm.DB
	service module.Service
}

func NewAccountHandler(db *gorm.DB, service module.Service) AccountHandler {
	return &accountHandler{
		db:      db,
		service: service,
	}
}

// GET /api/accounts/me
func (ctrl *accountHandler) ApiGetMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	account, err := ctrl.service.GetOne(module.GetOneDto{Id: accountId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAccountResponse(account))
}

// PUT /api/accounts/me
func (ctrl *accountHandler) ApiPutMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req module.PutMeRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	account, err := ctrl.service.UpdateOne(module.UpdateOneDto{
		Id:   accountId,
		Name: req.Name,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAccountResponse(account))
}

// DELETE /api/accounts/me
func (ctrl *accountHandler) ApiDeleteMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	if err := ctrl.service.DeleteOne(module.DeleteOneDto{Id: accountId}, ctrl.db); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}
