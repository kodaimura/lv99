package account_extended

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/helper"
	usecase "lv99/internal/usecase/account_extended"
)

// -----------------------------
// Handler Interface
// -----------------------------

type Handler interface {
	AdminGetWithProfile(c *gin.Context)
	AdminGetOneWithProfile(c *gin.Context)
	ApiGetAdminWithProfile(c *gin.Context)
}

type handler struct {
	usecase usecase.Usecase
}

func NewHandler(usecase usecase.Usecase) Handler {
	return &handler{
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// GET /api/admin/accounts/with-profile
func (h *handler) AdminGetWithProfile(c *gin.Context) {
	accounts, err := h.usecase.GetWithProfile(usecase.GetWithProfileDto{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountWithProfileResponseList(accounts))
}

// GET /api/admin/accounts/:account_id/with-profile
func (h *handler) AdminGetOneWithProfile(c *gin.Context) {
	var uri AccountUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	account, err := h.usecase.GetOneWithProfile(usecase.GetOneWithProfileDto{Id: uri.AccountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountWithProfileResponse(account))
}

// GET /api/accounts/admin/with-profile
func (h *handler) ApiGetAdminWithProfile(c *gin.Context) {
	account, err := h.usecase.GetAdminWithProfile()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAccountWithProfileResponse(account))
}
