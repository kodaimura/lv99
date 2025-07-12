package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
	module "lv99/internal/module/account_profile"
)

type AccountProfileHandler interface {
	ApiGetMe(c *gin.Context)
	ApiPutMe(c *gin.Context)
}

type accountProfileHandler struct {
	db      *gorm.DB
	service module.Service
}

func NewAccountProfileHandler(db *gorm.DB, service module.Service) AccountProfileHandler {
	return &accountProfileHandler{
		db:      db,
		service: service,
	}
}

// GET /api/accounts/me/profile
func (ctrl *accountProfileHandler) ApiGetMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	profile, err := ctrl.service.GetOne(module.GetOneDto{AccountId: accountId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAccountProfileResponse(profile))
}

// PUT /api/accounts/me/profile
func (ctrl *accountProfileHandler) ApiPutMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req module.PutMeRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	profile, err := ctrl.service.UpdateOne(module.UpdateOneDto{
		AccountId:   accountId,
		DisplayName: req.DisplayName,
		Bio:         req.Bio,
		AvatarURL:   req.AvatarURL,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAccountProfileResponse(profile))
}
