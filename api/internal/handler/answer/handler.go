package answer

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/helper"
	usecase "lv99/internal/usecase/answer"
)

// -----------------------------
// Handler Interface
// -----------------------------

type Handler interface {
	ApiGet(c *gin.Context)
	ApiPostOne(c *gin.Context)
	ApiGetOne(c *gin.Context)
	ApiPutOne(c *gin.Context)
	ApiDeleteOne(c *gin.Context)

	AdminGet(c *gin.Context)
	AdminGetOne(c *gin.Context)
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

// GET /api/answers?question_id=:question_id
func (h *handler) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req GetAnswersRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := h.usecase.Get(usecase.GetDto{
		QuestionId: req.QuestionId,
		AccountId:  accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponseList(answers))
}

// POST /api/answers
func (h *handler) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req PostAnswersRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := h.usecase.CreateOne(usecase.CreateOneDto{
		QuestionId: req.QuestionId,
		AccountId:  accountId,
		CodeDef:    req.CodeDef,
		CodeCall:   req.CodeCall,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, ToAnswerResponse(answer))
}

// GET /api/answers/:answer_id
func (h *handler) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	answer, err := h.usecase.GetOne(usecase.GetOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponse(answer))
}

// PUT /api/answers/:answer_id
func (h *handler) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	var req PutAnswerRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := h.usecase.UpdateOne(usecase.UpdateOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
		CodeDef:   req.CodeDef,
		CodeCall:  req.CodeCall,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponse(answer))
}

// DELETE /api/answers/:answer_id
func (h *handler) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := h.usecase.DeleteOne(usecase.DeleteOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}

// GET /api/admin/answers?account_id=:account_id&question_id=:question_id
func (h *handler) AdminGet(c *gin.Context) {
	var req AdminGetAnswersRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := h.usecase.Get(usecase.GetDto(req))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponseList(answers))
}

// GET /api/admin/answers/:answer_id
func (h *handler) AdminGetOne(c *gin.Context) {
	var uri AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	ans, err := h.usecase.GetOne(usecase.GetOneDto{
		Id: uri.AnswerId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponse(ans))
}
