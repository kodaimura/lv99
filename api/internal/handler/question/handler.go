package question

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/helper"
	usecase "lv99/internal/usecase/question"
)

// -----------------------------
// Handler Interface
// -----------------------------

type Handler interface {
	ApiGet(c *gin.Context)
	ApiGetOne(c *gin.Context)
	AdminGet(c *gin.Context)
	AdminGetOne(c *gin.Context)
	AdminPostOne(c *gin.Context)
	AdminPutOne(c *gin.Context)
	AdminDeleteOne(c *gin.Context)
	AdminRestoreOne(c *gin.Context)
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

// GET /api/questions
func (h *handler) ApiGet(c *gin.Context) {
	questions, err := h.usecase.Get(usecase.GetDto{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponseList(questions))
}

// GET /api/questions/:question_id
func (h *handler) ApiGetOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := h.usecase.GetOne(usecase.GetOneDto{Id: uri.QuestionId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponse(question))
}

// GET /api/admin/questions
func (h *handler) AdminGet(c *gin.Context) {
	questions, err := h.usecase.GetAll(usecase.GetAllDto{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponseList(questions))
}

// GET /api/admin/questions/:question_id
func (h *handler) AdminGetOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := h.usecase.GetOne(usecase.GetOneDto{Id: uri.QuestionId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponse(question))
}

// POST /api/admin/questions
func (h *handler) AdminPostOne(c *gin.Context) {
	var req PostQuestionRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := h.usecase.CreateOne(usecase.CreateOneDto(req))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, ToQuestionResponse(question))
}

// PUT /api/admin/questions/:question_id
func (h *handler) AdminPutOne(c *gin.Context) {
	var uri QuestionUri
	var req PutQuestionRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := h.usecase.UpdateOne(usecase.UpdateOneDto{
		Id:      uri.QuestionId,
		Title:   req.Title,
		Content: req.Content,
		Answer:  req.Answer,
		Level:   req.Level,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponse(question))
}

// DELETE /api/admin/questions/:question_id
func (h *handler) AdminDeleteOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := h.usecase.DeleteOne(usecase.DeleteOneDto{Id: uri.QuestionId}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}

// PATCH /api/admin/questions/:question_id
func (h *handler) AdminRestoreOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := h.usecase.RestoreOne(usecase.RestoreOneDto{Id: uri.QuestionId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
