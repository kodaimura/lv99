package answer_extended

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/helper"
	usecase "lv99/internal/usecase/answer_extended"
)

// -----------------------------
// Handler Interface
// -----------------------------

type Handler interface {
	ApiGetStatus(c *gin.Context)
	AdminGetStatus(c *gin.Context)
	AdminSearch(c *gin.Context)
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

// GET /api/answers/status
func (h *handler) ApiGetStatus(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	status, err := h.usecase.GetStatus(usecase.GetStatusDto{AccountId: accountId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerStatusResponseList(status))
}

// GET /api/admin/answers/status?account_id=:account_id
func (h *handler) AdminGetStatus(c *gin.Context) {
	var req GetStatusRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}
	status, err := h.usecase.GetStatus(usecase.GetStatusDto(req))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerStatusResponseList(status))
}

// GET /api/admin/answers/search?...
func (h *handler) AdminSearch(c *gin.Context) {
	var req SearchRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := h.usecase.Search(usecase.SearchDto(req))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerSearchResponseList(answers))
}
