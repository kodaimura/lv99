package comment

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/helper"
	usecase "lv99/internal/usecase/comment"
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

// GET /api/comments?answer_id=:answer_id
func (h *handler) ApiGet(c *gin.Context) {
	var req GetCommentsRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	comments, err := h.usecase.Get(usecase.GetDto{
		AnswerId: req.AnswerId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToCommentResponseList(comments))
}

// POST /api/comments
func (h *handler) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req PostOneRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := h.usecase.CreateOne(usecase.CreateOneDto{
		AnswerId:  req.AnswerId,
		AccountId: accountId,
		Content:   req.Content,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, ToCommentResponse(comment))
}

// GET /api/comments/:comment_id
func (h *handler) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri CommentUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	comment, err := h.usecase.GetOne(usecase.GetOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToCommentResponse(comment))
}

// PUT /api/comments/:comment_id
func (h *handler) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri CommentUri
	var req PutOneRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := h.usecase.UpdateOne(usecase.UpdateOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
		Content:   req.Content,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToCommentResponse(comment))
}

// DELETE /api/comments/:comment_id
func (h *handler) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri CommentUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := h.usecase.DeleteOne(usecase.DeleteOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
