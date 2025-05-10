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
	var req request.GetComment
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}

	comments, err := ctrl.commentService.Get(input.Comment{
		AnswerId: req.AnswerId,
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
	var req request.PostComment
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.commentService.CreateOne(input.Comment{
		AnswerId:  req.AnswerId,
		AccountId: accountId,
		Content:   req.Content,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, response.FromModelComment(comment))
}

// GET /api/comments/:id
func (ctrl *CommentController) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req request.CommentPK
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.commentService.GetOne(input.Comment{
		Id:        req.Id,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelComment(comment))
}

// PUT /api/comments/:id
func (ctrl *CommentController) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req request.PutComment
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	comment, err := ctrl.commentService.UpdateOne(input.Comment{
		Id:        req.Id,
		AccountId: accountId,
		Content:   req.Content,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelComment(comment))
}

// DELETE /api/comments/:id
func (ctrl *CommentController) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req request.CommentPK
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.commentService.DeleteOne(input.Comment{
		Id:        req.Id,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
