package controller

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/dto/input"
	"lv99/internal/dto/request"
	"lv99/internal/dto/response"
	"lv99/internal/helper"
	"lv99/internal/service"
)

type QuestionController struct {
	questionService service.QuestionService
}

func NewQuestionController(questionService service.QuestionService) *QuestionController {
	return &QuestionController{
		questionService: questionService,
	}
}

// GET /api/questions
func (ctrl *QuestionController) ApiGet(c *gin.Context) {
	questions, err := ctrl.questionService.Get(input.Question{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelQuestionList(questions))
}

// GET /api/questions/:question_id
func (ctrl *QuestionController) ApiGetOne(c *gin.Context) {
	var uri request.QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.questionService.GetOne(input.Question{Id: uri.QuestionId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelQuestion(question))
}

// GET /api/admin/questions
func (ctrl *QuestionController) AdminGet(c *gin.Context) {
	questions, err := ctrl.questionService.GetAll(input.Question{})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelQuestionList(questions))
}

// GET /api/admin/questions/:question_id
func (ctrl *QuestionController) AdminGetOne(c *gin.Context) {
	var uri request.QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.questionService.GetOne(input.Question{Id: uri.QuestionId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelQuestion(question))
}

// POST /api/admin/questions
func (ctrl *QuestionController) AdminPostOne(c *gin.Context) {
	var req request.QuestionBody
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.questionService.CreateOne(input.Question{
		Title:   req.Title,
		Content: req.Content,
		Answer:  req.Answer,
		Level:   req.Level,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, response.FromModelQuestion(question))
}

// PUT /api/admin/questions/:question_id
func (ctrl *QuestionController) AdminPutOne(c *gin.Context) {
	var uri request.QuestionUri
	var req request.QuestionBody
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.questionService.UpdateOne(input.Question{
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

	c.JSON(200, response.FromModelQuestion(question))
}

// DELETE /api/admin/questions/:question_id
func (ctrl *QuestionController) AdminDeleteOne(c *gin.Context) {
	var uri request.QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := ctrl.questionService.DeleteOne(input.Question{Id: uri.QuestionId}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}

// PATCH /api/admin/questions/:question_id
func (ctrl *QuestionController) AdminRestoreOne(c *gin.Context) {
	var uri request.QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.questionService.RestoreOne(input.Question{Id: uri.QuestionId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
