package question

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
)

type Controller interface {
	ApiGet(c *gin.Context)
	ApiGetOne(c *gin.Context)
	AdminGet(c *gin.Context)
	AdminGetOne(c *gin.Context)
	AdminPostOne(c *gin.Context)
	AdminPutOne(c *gin.Context)
	AdminDeleteOne(c *gin.Context)
	AdminRestoreOne(c *gin.Context)
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

// GET /api/questions
func (ctrl *controller) ApiGet(c *gin.Context) {
	questions, err := ctrl.service.Get(GetDto{}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponseList(questions))
}

// GET /api/questions/:question_id
func (ctrl *controller) ApiGetOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.service.GetOne(GetOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponse(question))
}

// GET /api/admin/questions
func (ctrl *controller) AdminGet(c *gin.Context) {
	questions, err := ctrl.service.GetAll(GetAllDto{}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponseList(questions))
}

// GET /api/admin/questions/:question_id
func (ctrl *controller) AdminGetOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.service.GetOne(GetOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponse(question))
}

// POST /api/admin/questions
func (ctrl *controller) AdminPostOne(c *gin.Context) {
	var req PostOneRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.service.CreateOne(CreateOneDto(req), ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, ToQuestionResponse(question))
}

// PUT /api/admin/questions/:question_id
func (ctrl *controller) AdminPutOne(c *gin.Context) {
	var uri QuestionUri
	var req PutOneRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.service.UpdateOne(UpdateOneDto{
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

	c.JSON(200, ToQuestionResponse(question))
}

// DELETE /api/admin/questions/:question_id
func (ctrl *controller) AdminDeleteOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := ctrl.service.DeleteOne(DeleteOneDto{Id: uri.QuestionId}, ctrl.db); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}

// PATCH /api/admin/questions/:question_id
func (ctrl *controller) AdminRestoreOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.service.RestoreOne(RestoreOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
