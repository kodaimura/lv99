package controller

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/dto/input"
	"lv99/internal/dto/request"
	"lv99/internal/dto/response"
	"lv99/internal/helper"
	"lv99/internal/service"
)

type AnswerController struct {
	answerService service.AnswerService
}

func NewAnswerController(answerService service.AnswerService) *AnswerController {
	return &AnswerController{
		answerService: answerService,
	}
}

// GET /api/answers/:question_id/answers
func (ctrl *AnswerController) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.QuestionPK
	answers, err := ctrl.answerService.Get(input.Answer{
		QuestionId: uri.QuestionId, 
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelAnswerList(answers))
}

// POST /api/answers/:question_id/answers
func (ctrl *AnswerController) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri request.QuestionPK
	var req request.Answer
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.answerService.CreateOne(input.Answer{
		QuestionId:   uri.QuestionId,
		AccountId: accountId,
		CodeDef: req.CodeDef,
		CodeCall: req.CodeCall,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, response.FromModelAnswer(answer))
}