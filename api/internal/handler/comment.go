package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
	commentModule "lv99/internal/module/comment"
	usecase "lv99/internal/usecase/comment"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type CommentResponse struct {
	Id        int            `json:"id"`
	AnswerId  int            `json:"answer_id"`
	AccountId int            `json:"account_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func ToCommentResponse(m commentModule.Comment) CommentResponse {
	return CommentResponse(m)
}

func ToCommentResponseList(models []commentModule.Comment) []CommentResponse {
	res := make([]CommentResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToCommentResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type CommentUri struct {
	CommentId int `uri:"comment_id" binding:"required"`
}

type GetCommentsRequest struct {
	AnswerId int `form:"answer_id"`
}

type PostCommentRequest struct {
	AnswerId int    `json:"answer_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type PutCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// -----------------------------
// Handler Interface
// -----------------------------

type CommentHandler interface {
	ApiGet(c *gin.Context)
	ApiPostOne(c *gin.Context)
	ApiGetOne(c *gin.Context)
	ApiPutOne(c *gin.Context)
	ApiDeleteOne(c *gin.Context)
}

type commentHandler struct {
	db      *gorm.DB
	usecase usecase.Usecase
}

func NewCommentHandler(db *gorm.DB, usecase usecase.Usecase) CommentHandler {
	return &commentHandler{
		db:      db,
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// GET /api/comments?answer_id=:answer_id
func (ctrl *commentHandler) ApiGet(c *gin.Context) {
	var req GetCommentsRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	comments, err := ctrl.usecase.Get(usecase.GetDto{
		AnswerId: req.AnswerId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToCommentResponseList(comments))
}

// POST /api/comments
func (ctrl *commentHandler) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req PostCommentRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.usecase.CreateOne(usecase.CreateOneDto{
		AnswerId:  req.AnswerId,
		AccountId: accountId,
		Content:   req.Content,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, ToCommentResponse(comment))
}

// GET /api/comments/:comment_id
func (ctrl *commentHandler) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri CommentUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.usecase.GetOne(usecase.GetOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToCommentResponse(comment))
}

// PUT /api/comments/:comment_id
func (ctrl *commentHandler) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri CommentUri
	var req PutCommentRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.usecase.UpdateOne(usecase.UpdateOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
		Content:   req.Content,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToCommentResponse(comment))
}

// DELETE /api/comments/:comment_id
func (ctrl *commentHandler) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri CommentUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.usecase.DeleteOne(usecase.DeleteOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
