package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
	module "lv99/internal/module/question"
)

type QuestionHandler interface {
	ApiGet(c *gin.Context)
	ApiGetOne(c *gin.Context)
	AdminGet(c *gin.Context)
	AdminGetOne(c *gin.Context)
	AdminPostOne(c *gin.Context)
	AdminPutOne(c *gin.Context)
	AdminDeleteOne(c *gin.Context)
	AdminRestoreOne(c *gin.Context)
}

type questionHandler struct {
	db      *gorm.DB
	service module.Service
}

func NewQuestionHandler(db *gorm.DB, service module.Service) QuestionHandler {
	return &questionHandler{
		db:      db,
		service: service,
	}
}

// GET /api/questions
func (ctrl *questionHandler) ApiGet(c *gin.Context) {
	questions, err := ctrl.service.Get(module.GetDto{}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToQuestionResponseList(questions))
}

// GET /api/questions/:question_id
func (ctrl *questionHandler) ApiGetOne(c *gin.Context) {
	var uri module.QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.service.GetOne(module.GetOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToQuestionResponse(question))
}

// GET /api/admin/questions
func (ctrl *questionHandler) AdminGet(c *gin.Context) {
	questions, err := ctrl.service.GetAll(module.GetAllDto{}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToQuestionResponseList(questions))
}

// GET /api/admin/questions/:question_id
func (ctrl *questionHandler) AdminGetOne(c *gin.Context) {
	var uri module.QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.service.GetOne(module.GetOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToQuestionResponse(question))
}

// POST /api/admin/questions
func (ctrl *questionHandler) AdminPostOne(c *gin.Context) {
	var req module.PostOneRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.service.CreateOne(module.CreateOneDto(req), ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, module.ToQuestionResponse(question))
}

// PUT /api/admin/questions/:question_id
func (ctrl *questionHandler) AdminPutOne(c *gin.Context) {
	var uri module.QuestionUri
	var req module.PutOneRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.service.UpdateOne(module.UpdateOneDto{
		Id:      uri.QuestionId,
		Title:   req.Title,
		Content: req.Content,
		Answer:  req.Answer,
		Level:   req.Level,
	}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, module.ToQuestionResponse(question))
}

// DELETE /api/admin/questions/:question_id
func (ctrl *questionHandler) AdminDeleteOne(c *gin.Context) {
	var uri module.QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := ctrl.service.DeleteOne(module.DeleteOneDto{Id: uri.QuestionId}, ctrl.db); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}

// PATCH /api/admin/questions/:question_id
func (ctrl *questionHandler) AdminRestoreOne(c *gin.Context) {
	var uri module.QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.service.RestoreOne(module.RestoreOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
