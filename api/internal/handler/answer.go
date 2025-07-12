package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
	module "lv99/internal/module/answer"
)

type AnswerHandler interface {
	ApiGet(c *gin.Context)
	ApiPostOne(c *gin.Context)
	ApiGetOne(c *gin.Context)
	ApiPutOne(c *gin.Context)
	ApiDeleteOne(c *gin.Context)

	AdminGet(c *gin.Context)
	AdminGetOne(c *gin.Context)
}

type answerHandler struct {
	db      *gorm.DB
	service module.Service
}

func NewAnswerHandler(db *gorm.DB, service module.Service) AnswerHandler {
	return &answerHandler{
		db:      db,
		service: service,
	}
}

// GET /api/answers?question_id=:question_id
func (ctrl *answerHandler) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req module.GetRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := ctrl.service.Get(module.GetDto{
		QuestionId: req.QuestionId,
		AccountId:  accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAnswerResponseList(answers))
}

// POST /api/answers
func (ctrl *answerHandler) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req module.PostOneRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	ans, err := ctrl.service.CreateOne(module.CreateOneDto{
		QuestionId: req.QuestionId,
		AccountId:  accountId,
		CodeDef:    req.CodeDef,
		CodeCall:   req.CodeCall,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, module.ToAnswerResponse(ans))
}

// GET /api/answers/:answer_id
func (ctrl *answerHandler) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri module.AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	ans, err := ctrl.service.GetOne(module.GetOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAnswerResponse(ans))
}

// PUT /api/answers/:answer_id
func (ctrl *answerHandler) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri module.AnswerUri
	var req module.PutOneRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	ans, err := ctrl.service.UpdateOne(module.UpdateOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
		CodeDef:   req.CodeDef,
		CodeCall:  req.CodeCall,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAnswerResponse(ans))
}

// DELETE /api/answers/:answer_id
func (ctrl *answerHandler) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri module.AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.service.DeleteOne(module.DeleteOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}

// GET /api/admin/answers?account_id=:account_id&question_id=:question_id
func (ctrl *answerHandler) AdminGet(c *gin.Context) {
	var req module.AdminGetRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := ctrl.service.Get(module.GetDto(req), ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAnswerResponseList(answers))
}

// GET /api/admin/answers/:answer_id
func (ctrl *answerHandler) AdminGetOne(c *gin.Context) {
	var uri module.AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	ans, err := ctrl.service.GetOne(module.GetOneDto{
		Id: uri.AnswerId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToAnswerResponse(ans))
}
