package controller

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/dto/input"
	"lv99/internal/dto/request"
	"lv99/internal/dto/response"
	"lv99/internal/helper"
	"lv99/internal/service"
)

type AccountProfileController struct {
	accountProfileService service.AccountProfileService
}

func NewAccountProfileController(accountProfileService service.AccountProfileService) *AccountProfileController {
	return &AccountProfileController{
		accountProfileService: accountProfileService,
	}
}

// GET /api/accounts/me/profile
func (ctrl *AccountProfileController) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	profile, err := ctrl.accountProfileService.GetOne(input.AccountProfile{AccountId: accountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelAccountProfile(profile))
}

// PUT /api/accounts/me/profile
func (ctrl *AccountProfileController) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req request.PutAccountProfile
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	profile, err := ctrl.accountProfileService.UpdateOne(input.AccountProfile{
		AccountId:   accountId,
		DisplayName: req.DisplayName,
		Bio:         req.Bio,
		AvatarURL:   req.AvatarURL,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelAccountProfile(profile))
}
