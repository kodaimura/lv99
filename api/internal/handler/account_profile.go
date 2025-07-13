package handler

import (
	"time"

	"github.com/gin-gonic/gin"

	"lv99/internal/helper"
	profileModule "lv99/internal/module/account_profile"
	usecase "lv99/internal/usecase/account_profile"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type AccountProfileResponse struct {
	AccountId   int       `json:"account_id"`
	DisplayName string    `json:"display_name"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatar_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToAccountProfileResponse(m profileModule.AccountProfile) AccountProfileResponse {
	return AccountProfileResponse{
		AccountId:   m.AccountId,
		DisplayName: m.DisplayName,
		Bio:         m.Bio,
		AvatarURL:   m.AvatarURL,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func ToAccountProfileResponseList(models []profileModule.AccountProfile) []AccountProfileResponse {
	res := make([]AccountProfileResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAccountProfileResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type PutAccountMeProfileRequest struct {
	DisplayName string `json:"display_name" binding:"required"`
	Bio         string `json:"bio" binding:"omitempty"`
	AvatarURL   string `json:"avatar_url" binding:"omitempty,url"`
}

// -----------------------------
// Handler Interface
// -----------------------------

type AccountProfileHandler interface {
	ApiGetMe(c *gin.Context)
	ApiPutMe(c *gin.Context)
}

type accountProfileHandler struct {
	usecase usecase.Usecase
}

func NewAccountProfileHandler(usecase usecase.Usecase) AccountProfileHandler {
	return &accountProfileHandler{
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// GET /api/accounts/me/profile
func (ctrl *accountProfileHandler) ApiGetMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	profile, err := ctrl.usecase.GetOne(usecase.GetOneDto{AccountId: accountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountProfileResponse(profile))
}

// PUT /api/accounts/me/profile
func (ctrl *accountProfileHandler) ApiPutMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req PutAccountMeProfileRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	profile, err := ctrl.usecase.UpdateOne(usecase.UpdateOneDto{
		AccountId:   accountId,
		DisplayName: req.DisplayName,
		Bio:         req.Bio,
		AvatarURL:   req.AvatarURL,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountProfileResponse(profile))
}
