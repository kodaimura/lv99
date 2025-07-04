package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/config"
	"lv99/internal/core"
	"lv99/internal/helper"
	profile "lv99/internal/module/account_profile"
)

type Controller interface {
	ApiSignup(c *gin.Context)
	ApiLogin(c *gin.Context)
	ApiRefresh(c *gin.Context)
	ApiLogout(c *gin.Context)

	ApiPutMePassword(c *gin.Context)
}

type controller struct {
	db                    *gorm.DB
	service               Service
	accountProfileService profile.Service
}

func NewController(db *gorm.DB, service Service, accountProfileService profile.Service) Controller {
	return &controller{
		db:                    db,
		service:               service,
		accountProfileService: accountProfileService,
	}
}

// POST /api/accounts/signup
func (ctrl *controller) ApiSignup(c *gin.Context) {
	var req SignupRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.db.Transaction(func(tx *gorm.DB) error {
		acct, err := ctrl.service.Signup(SignupDto(req), tx)
		if err != nil {
			return err
		}
		_, err = ctrl.accountProfileService.CreateOne(profile.CreateOneDto{
			AccountId:   acct.Id,
			DisplayName: acct.Name,
			Bio:         "",
			AvatarURL:   "",
		}, tx)
		return err
	})

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, gin.H{})
}

// POST /api/accounts/login
func (ctrl *controller) ApiLogin(c *gin.Context) {
	var req LoginRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	acct, err := ctrl.service.Login(LoginDto(req), ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	accessToken, err := core.Auth.CreateAccessToken(core.AuthPayload{
		AccountId:   acct.Id,
		AccountName: acct.Name,
		AccountRole: acct.Role,
	})
	if err != nil {
		c.Error(err)
		return
	}

	refreshToken, err := core.Auth.CreateRefreshToken(core.AuthPayload{
		AccountId:   acct.Id,
		AccountName: acct.Name,
		AccountRole: acct.Role,
	})
	if err != nil {
		c.Error(err)
		return
	}

	helper.SetAccessTokenCookie(c, accessToken)
	helper.SetRefreshTokenCookie(c, refreshToken)

	core.Logger.Info("account login: id=%d name=%s", acct.Id, acct.Name)

	c.JSON(200, LoginResponse{
		AccountId:        acct.Id,
		AccountRole:      acct.Role,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  config.AccessTokenExpiresSeconds,
		RefreshExpiresIn: config.RefreshTokenExpiresSeconds,
	})
}

// POST /api/accounts/refresh
func (ctrl *controller) ApiRefresh(c *gin.Context) {
	refreshToken := helper.GetRefreshToken(c)

	payload, err := core.Auth.VerifyRefreshToken(refreshToken)
	if err != nil {
		c.Error(core.NewAppError("invalid or expired refresh token", core.ErrCodeUnauthorized))
		return
	}

	accessToken, err := core.Auth.CreateAccessToken(core.AuthPayload{
		AccountId:   payload.AccountId,
		AccountName: payload.AccountName,
		AccountRole: payload.AccountRole,
	})
	if err != nil {
		c.Error(err)
		return
	}

	helper.SetAccessTokenCookie(c, accessToken)

	core.Logger.Info("access token refreshed: id=%d name=%s", payload.AccountId, payload.AccountName)

	c.JSON(200, RefreshResponse{
		AccessToken: accessToken,
		ExpiresIn:   config.AccessTokenExpiresSeconds,
	})
}

// POST /api/accounts/logout
func (ctrl *controller) ApiLogout(c *gin.Context) {
	core.Auth.RevokeRefreshToken(helper.GetRefreshToken(c))
	helper.SetAccessTokenCookie(c, "")
	helper.SetRefreshTokenCookie(c, "")
	c.JSON(200, gin.H{})
}

// PUT /api/accounts/me/password
func (ctrl *controller) ApiPutMePassword(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req PutMePasswordRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.service.UpdatePassword(UpdatePasswordDto{
		Id:          accountId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{})
}
