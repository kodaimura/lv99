package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/config"
	"lv99/internal/core"
	"lv99/internal/helper"
	"lv99/internal/module/auth"
	usecase "lv99/internal/usecase/auth"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type LoginResponse struct {
	AccountId        int    `json:"account_id"`
	AccountRole      int    `json:"account_role"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	AccessExpiresIn  int    `json:"access_expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// -----------------------------
// DTO（Request）
// -----------------------------

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PutMePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// -----------------------------
// Handler Interface
// -----------------------------

type AuthHandler interface {
	ApiSignup(c *gin.Context)
	ApiLogin(c *gin.Context)
	ApiRefresh(c *gin.Context)
	ApiLogout(c *gin.Context)

	ApiPutMePassword(c *gin.Context)
}

type authHandler struct {
	db      *gorm.DB
	usecase usecase.Usecase
}

func NewAuthHandler(db *gorm.DB, usecase usecase.Usecase) AuthHandler {
	return &authHandler{
		db:      db,
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// POST /api/accounts/signup
func (ctrl *authHandler) ApiSignup(c *gin.Context) {
	var req auth.SignupRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	_, err := ctrl.usecase.Signup(usecase.SignupDto(req), ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, gin.H{})
}

// POST /api/accounts/login
func (ctrl *authHandler) ApiLogin(c *gin.Context) {
	var req auth.LoginRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	acct, accessToken, refreshToken, err := ctrl.usecase.Login(usecase.LoginDto(req), ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	helper.SetAccessTokenCookie(c, accessToken)
	helper.SetRefreshTokenCookie(c, refreshToken)
	core.Logger.Info("account login: id=%d name=%s", acct.Id, acct.Name)

	c.JSON(200, auth.LoginResponse{
		AccountId:        acct.Id,
		AccountRole:      acct.Role,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  config.AccessTokenExpiresSeconds,
		RefreshExpiresIn: config.RefreshTokenExpiresSeconds,
	})
}

// POST /api/accounts/refresh
func (ctrl *authHandler) ApiRefresh(c *gin.Context) {
	refreshToken := helper.GetRefreshToken(c)

	payload, accessToken, err := ctrl.usecase.Refresh(refreshToken, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	helper.SetAccessTokenCookie(c, accessToken)
	core.Logger.Info("access token refreshed: id=%d name=%s", payload.AccountId, payload.AccountName)

	c.JSON(200, auth.RefreshResponse{
		AccessToken: accessToken,
		ExpiresIn:   config.AccessTokenExpiresSeconds,
	})
}

// POST /api/accounts/logout
func (ctrl *authHandler) ApiLogout(c *gin.Context) {
	core.Auth.RevokeRefreshToken(helper.GetRefreshToken(c))
	helper.SetAccessTokenCookie(c, "")
	helper.SetRefreshTokenCookie(c, "")
	c.JSON(200, gin.H{})
}

// PUT /api/accounts/me/password
func (ctrl *authHandler) ApiPutMePassword(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req auth.PutMePasswordRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.usecase.UpdatePassword(usecase.UpdatePasswordDto{
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
