package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
	module "lv99/internal/module/comment"
)

type CommentHandler interface {
	ApiGet(c *gin.Context)
	ApiPostOne(c *gin.Context)
	ApiGetOne(c *gin.Context)
	ApiPutOne(c *gin.Context)
	ApiDeleteOne(c *gin.Context)
}

type commentHandler struct {
	db      *gorm.DB
	service module.Service
}

func NewCommentHandler(db *gorm.DB, service module.Service) CommentHandler {
	return &commentHandler{
		db:      db,
		service: service,
	}
}

// GET /api/comments?answer_id=:answer_id
func (ctrl *commentHandler) ApiGet(c *gin.Context) {
	var req module.GetRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	comments, err := ctrl.service.Get(module.GetDto{
		AnswerId: req.AnswerId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToCommentResponseList(comments))
}

// POST /api/comments
func (ctrl *commentHandler) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req module.PostOneRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.service.CreateOne(module.CreateOneDto{
		AnswerId:  req.AnswerId,
		AccountId: accountId,
		Content:   req.Content,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, module.ToCommentResponse(comment))
}

// GET /api/comments/:comment_id
func (ctrl *commentHandler) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri module.CommentUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.service.GetOne(module.GetOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToCommentResponse(comment))
}

// PUT /api/comments/:comment_id
func (ctrl *commentHandler) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri module.CommentUri
	var req module.PutOneRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.service.UpdateOne(module.UpdateOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
		Content:   req.Content,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToCommentResponse(comment))
}

// DELETE /api/comments/:comment_id
func (ctrl *commentHandler) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri module.CommentUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.service.DeleteOne(module.DeleteOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
