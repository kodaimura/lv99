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

// POST /api/questions
func (ctrl *QuestionController) ApiPostOne(c *gin.Context) {
	var req request.Question
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.questionService.CreateOne(input.Question{
		QuestionTitle:   req.QuestionTitle,
		QuestionContent: req.QuestionContent,
		QuestionAnswer:  req.QuestionAnswer,
		QuestionLevel:   req.QuestionLevel,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, response.FromModelQuestion(question))
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
	var uri request.QuestionPK
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.questionService.GetOne(input.Question{QuestionId: uri.QuestionId})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelQuestion(question))
}

// PUT /api/questions/:question_id
func (ctrl *QuestionController) ApiPutOne(c *gin.Context) {
	var uri request.QuestionPK
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	var req request.Question
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.questionService.UpdateOne(input.Question{
		QuestionId:      uri.QuestionId,
		QuestionTitle:   req.QuestionTitle,
		QuestionContent: req.QuestionContent,
		QuestionAnswer:  req.QuestionAnswer,
		QuestionLevel:   req.QuestionLevel,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelQuestion(question))
}

// DELETE /api/questions/:question_id
func (ctrl *QuestionController) ApiDeleteOne(c *gin.Context) {
	var uri request.QuestionPK
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := ctrl.questionService.DeleteOne(input.Question{QuestionId: uri.QuestionId}); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}
