package answer

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

// GET /api/questions/:question_id/answers
func (ctrl *controller) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	answers, err := ctrl.service.Get(GetDto{
		QuestionId: uri.QuestionId,
		AccountId:  accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponseList(answers))
}

// POST /api/questions/:question_id/answers
func (ctrl *controller) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri QuestionUri
	var req PostOneRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.service.CreateOne(CreateOneDto{
		QuestionId: uri.QuestionId,
		AccountId:  accountId,
		CodeDef:    req.CodeDef,
		CodeCall:   req.CodeCall,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, ToAnswerResponse(answer))
}

// GET /api/answers/:answer_id
func (ctrl *controller) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.service.GetOne(GetOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponse(answer))
}

// PUT /api/answers/:answer_id
func (ctrl *controller) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	var req PutOneRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.service.UpdateOne(UpdateOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
		CodeDef:   req.CodeDef,
		CodeCall:  req.CodeCall,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponse(answer))
}

// DELETE /api/answers/:answer_id
func (ctrl *controller) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.service.DeleteOne(DeleteOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
