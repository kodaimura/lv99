package comment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Controller interface {
	ApiGet(c *gin.Context)
	ApiPostOne(c *gin.Context)
	ApiGetOne(c *gin.Context)
	ApiPutOne(c *gin.Context)
	ApiDeleteOne(c *gin.Context)
}

type controller struct {
	db      *gorm.DB
	service Service
}

func NewController(db *gorm.DB, service Service) Controller {
	return &controller{
		db:      db,
		service: service,
	}
}

// GET /api/comments?answer_id=:answer_id
func (ctrl *controller) ApiGet(c *gin.Context) {
	var req GetRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	comments, err := ctrl.service.Get(GetDto{
		AnswerId: req.AnswerId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToCommentResponseList(comments))
}

// POST /api/comments
func (ctrl *controller) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req PostOneRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.service.CreateOne(CreateOneDto{
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
func (ctrl *controller) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri CommentUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.service.GetOne(GetOneDto{
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
func (ctrl *controller) ApiPutOne(c *gin.Context) {
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

	comment, err := ctrl.service.UpdateOne(UpdateOneDto{
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
func (ctrl *controller) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri CommentUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.service.DeleteOne(DeleteOneDto{
		Id:        uri.CommentId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
