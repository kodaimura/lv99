package controller

import (
	"github.com/gin-gonic/gin"

	"lv99/config"
	"lv99/internal/core"
	"lv99/internal/dto/input"
	"lv99/internal/dto/request"
	"lv99/internal/dto/response"
	"lv99/internal/helper"
	"lv99/internal/service"
)

type AdminController struct {
	adminService service.AdminService
}

func NewAdminController(adminService service.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

// POST /api/admins/signup
func (ctrl *AdminController) ApiSignup(c *gin.Context) {
	var req request.AdminSignup
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	admin, err := ctrl.adminService.Signup(input.AdminSignup{
		AdminName:     req.AdminName,
		AdminPassword: req.AdminPassword,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.Admin{
		AdminId:   admin.AdminId,
		AdminName: admin.AdminName,
		CreatedAt:   admin.CreatedAt,
		UpdatedAt:   admin.UpdatedAt,
	})
}

// POST /api/admins/login
func (ctrl *AdminController) ApiLogin(c *gin.Context) {
	var req request.AdminLogin
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	admin, err := ctrl.adminService.Login(input.AdminLogin{
		AdminName:     req.AdminName,
		AdminPassword: req.AdminPassword,
	})
	if err != nil {
		c.Error(err)
		return
	}

	accessToken, err := core.Auth.CreateAccessToken(core.AuthPayload{
		AccountId:   admin.AdminId,
		AccountName: admin.AdminName,
		AccountRole: 0,
	})
	if err != nil {
		c.Error(err)
		return
	}

	refreshToken, err := core.Auth.CreateRefreshToken(core.AuthPayload{
		AccountId:   admin.AdminId,
		AccountName: admin.AdminName,
		AccountRole: 0,
	})
	if err != nil {
		c.Error(err)
		return
	}

	helper.SetAccessTokenCookie(c, accessToken)
	helper.SetRefreshTokenCookie(c, refreshToken)

	core.Logger.Info("admin login: id=%d name=%s", admin.AdminId, admin.AdminName)

	c.JSON(200, response.AdminLogin{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresIn:  config.AccessTokenExpiresSeconds,
		RefreshExpiresIn: config.RefreshTokenExpiresSeconds,
		Admin: response.Admin{
			AdminId:   admin.AdminId,
			AdminName: admin.AdminName,
			CreatedAt:   admin.CreatedAt,
			UpdatedAt:   admin.UpdatedAt,
		},
	})
}

// POST /api/admins/refresh
func (ctrl *AdminController) ApiRefresh(c *gin.Context) {
	refreshToken := helper.GetRefreshToken(c)

	payload, err := core.Auth.VerifyRefreshToken(refreshToken)
	if err != nil {
		c.Error(core.NewAppError("invalid or expired refresh token", core.ErrCodeUnauthorized))
		return
	}

	accessToken, err := core.Auth.CreateAccessToken(core.AuthPayload{
		AccountId:   payload.AccountId,
		AccountName: payload.AccountName,
		AccountRole: 0,
	})
	if err != nil {
		c.Error(err)
		return
	}

	helper.SetAccessTokenCookie(c, accessToken)

	c.JSON(200, response.Refresh{
		AccessToken: accessToken,
		ExpiresIn:   config.AccessTokenExpiresSeconds,
	})
}

// POST /api/admins/logout
func (ctrl *AdminController) ApiLogout(c *gin.Context) {
	core.Auth.RevokeRefreshToken(helper.GetRefreshToken(c))
	helper.SetAccessTokenCookie(c, "")
	helper.SetRefreshTokenCookie(c, "")
	c.JSON(200, gin.H{})
}

// GET /api/admins/me
func (ctrl *AdminController) ApiGetOne(c *gin.Context) {
	adminId := helper.GetAccountId(c)
	admin, err := ctrl.adminService.GetOne(input.Admin{AdminId: adminId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.Admin{
		AdminId:   admin.AdminId,
		AdminName: admin.AdminName,
		CreatedAt:   admin.CreatedAt,
		UpdatedAt:   admin.UpdatedAt,
	})
}

// PUT /api/admins/me/password
func (ctrl *AdminController) ApiPutPassword(c *gin.Context) {
	adminName := helper.GetAccountName(c)

	var req request.PutAdminPassword
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	admin, err := ctrl.adminService.Login(input.AdminLogin{
		AdminName:     adminName,
		AdminPassword: req.OldAdminPassword,
	})
	if err != nil {
		c.Error(err)
		return
	}

	_, err = ctrl.adminService.UpdateOne(input.Admin{
		AdminId:       admin.AdminId,
		AdminPassword: req.NewAdminPassword,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{})
}

// PUT /api/admins/me
func (ctrl *AdminController) ApiPutOne(c *gin.Context) {
	adminId := helper.GetAccountId(c)

	var req request.PutAdmin
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	admin, err := ctrl.adminService.UpdateOne(input.Admin{
		AdminId:   adminId,
		AdminName: req.AdminName,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.Admin{
		AdminId:   admin.AdminId,
		AdminName: admin.AdminName,
		CreatedAt:   admin.CreatedAt,
		UpdatedAt:   admin.UpdatedAt,
	})
}

// DELETE /api/admins/me
func (ctrl *AdminController) ApiDeleteOne(c *gin.Context) {
	adminId := helper.GetAccountId(c)
	if err := ctrl.adminService.DeleteOne(input.Admin{AdminId: adminId}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{})
}
