package account_profile

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Controller interface {
	ApiGetMe(c *gin.Context)
	ApiPutMe(c *gin.Context)
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

// GET /api/accounts/me/profile
func (ctrl *controller) ApiGetMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	profile, err := ctrl.service.GetOne(GetOneDto{AccountId: accountId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountProfileResponse(profile))
}

// PUT /api/accounts/me/profile
func (ctrl *controller) ApiPutMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req PutMeRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	profile, err := ctrl.service.UpdateOne(UpdateOneDto{
		AccountId:   accountId,
		DisplayName: req.DisplayName,
		Bio:         req.Bio,
		AvatarURL:   req.AvatarURL,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountProfileResponse(profile))
}
