package handler

import (
	"time"

	"github.com/gin-gonic/gin"

	"lv99/internal/helper"
	accountModule "lv99/internal/module/account"
	usecase "lv99/internal/usecase/account"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type AccountResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccountResponse(m accountModule.Account) AccountResponse {
	return AccountResponse{
		Id:        m.Id,
		Name:      m.Name,
		Role:      m.Role,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToAccountResponseList(models []accountModule.Account) []AccountResponse {
	res := make([]AccountResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAccountResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type PutAccountMeRequest struct {
	Name string `json:"name" binding:"required"`
}

// -----------------------------
// Handler Interface
// -----------------------------

type AccountHandler interface {
	ApiGetMe(c *gin.Context)
	ApiPutMe(c *gin.Context)
	ApiDeleteMe(c *gin.Context)
}

type accountHandler struct {
	usecase usecase.Usecase
}

func NewAccountHandler(usecase usecase.Usecase) AccountHandler {
	return &accountHandler{
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// GET /api/accounts/me
func (ctrl *accountHandler) ApiGetMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	account, err := ctrl.usecase.GetOne(usecase.GetOneDto{Id: accountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountResponse(account))
}

// PUT /api/accounts/me
func (ctrl *accountHandler) ApiPutMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)

	var req PutAccountMeRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	account, err := ctrl.usecase.UpdateOne(usecase.UpdateOneDto{
		Id:   accountId,
		Name: req.Name,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountResponse(account))
}

// DELETE /api/accounts/me
func (ctrl *accountHandler) ApiDeleteMe(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	if err := ctrl.usecase.DeleteOne(usecase.DeleteOneDto{Id: accountId}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}
