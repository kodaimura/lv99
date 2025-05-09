package controller

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/dto/input"
	"lv99/internal/dto/request"
	"lv99/internal/dto/response"
	"lv99/internal/helper"
	"lv99/internal/service"
)

type CommentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

// GET /api/answers/:answer_id/comments
func (ctrl *CommentController) ApiGet(c *gin.Context) {
	var uri request.AnswerPK
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	comments, err := ctrl.commentService.Get(input.Comment{
		AnswerId: uri.AnswerId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelCommentList(comments))
}

// POST /api/answers/:answer_id/comments
func (ctrl *CommentController) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.AnswerPK
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	var req request.Comment
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.commentService.CreateOne(input.Comment{
		AnswerId:       uri.AnswerId,
		AccountId:      accountId,
		CommentContent: req.CommentContent,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, response.FromModelComment(comment))
}

// PUT /api/answers/:answer_id/comments/:comment_id
func (ctrl *CommentController) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.PutComment
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	var req request.Comment
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.commentService.UpdateOne(input.Comment{
		AnswerId:       uri.AnswerId,
		CommentId:      uri.CommentId,
		AccountId:      accountId,
		CommentContent: req.CommentContent,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelComment(comment))
}

// DELETE /api/answers/:answer_id/comments/:comment_id
func (ctrl *CommentController) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.DeleteComment
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.commentService.DeleteOne(input.Comment{
		AnswerId:  uri.AnswerId,
		CommentId: uri.CommentId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
