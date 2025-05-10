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

// GET /api/questions/:question_id/answers
func (ctrl *AnswerController) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req request.QuestionUri
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := ctrl.answerService.Get(input.Answer{
		QuestionId: req.QuestionId,
		AccountId:  accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelAnswerList(answers))
}

// POST /api/questions/:question_id/answers
func (ctrl *AnswerController) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req request.PostAnswer
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.answerService.CreateOne(input.Answer{
		QuestionId: req.QuestionId,
		AccountId:  accountId,
		CodeDef:    req.CodeDef,
		CodeCall:   req.CodeCall,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, response.FromModelAnswer(answer))
}

// GET /api/answers/:answer_id
func (ctrl *AnswerController) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req request.AnswerUri
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.answerService.GetOne(input.Answer{
		Id: req.AnswerId,
		AccountId:  accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelAnswer(answer))
}

// PUT /api/answers/:answer_id
func (ctrl *AnswerController) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req request.PutAnswer
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.answerService.UpdateOne(input.Answer{
		Id:   req.AnswerId,
		AccountId:  accountId,
		CodeDef:    req.CodeDef,
		CodeCall:   req.CodeCall,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, response.FromModelAnswer(answer))
}

// DELETE /api/answers/:answer_id
func (ctrl *AnswerController) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req request.AnswerUri
	if err := helper.BindUri(c, &req); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.answerService.DeleteOne(input.Answer{
		Id:   req.AnswerId,
		AccountId:  accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
