package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
	questionModule "lv99/internal/module/question"
	usecase "lv99/internal/usecase/question"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type QuestionResponse struct {
	Id        int            `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Answer    string         `json:"answer"`
	Level     int            `json:"level"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func ToQuestionResponse(m questionModule.Question) QuestionResponse {
	return QuestionResponse(m)
}

func ToQuestionResponseList(models []questionModule.Question) []QuestionResponse {
	res := make([]QuestionResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToQuestionResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type QuestionUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
}

type PostQuestionRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Level   int    `json:"level" binding:"required,min=1"`
}

type PutQuestionRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Level   int    `json:"level" binding:"required,min=1"`
}

// -----------------------------
// Handler Interface
// -----------------------------

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
	usecase usecase.Usecase
}

func NewQuestionHandler(db *gorm.DB, usecase usecase.Usecase) QuestionHandler {
	return &questionHandler{
		db:      db,
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// GET /api/questions
func (ctrl *questionHandler) ApiGet(c *gin.Context) {
	questions, err := ctrl.usecase.Get(usecase.GetDto{}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponseList(questions))
}

// GET /api/questions/:question_id
func (ctrl *questionHandler) ApiGetOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.usecase.GetOne(usecase.GetOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponse(question))
}

// GET /api/admin/questions
func (ctrl *questionHandler) AdminGet(c *gin.Context) {
	questions, err := ctrl.usecase.GetAll(usecase.GetAllDto{}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponseList(questions))
}

// GET /api/admin/questions/:question_id
func (ctrl *questionHandler) AdminGetOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	question, err := ctrl.usecase.GetOne(usecase.GetOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToQuestionResponse(question))
}

// POST /api/admin/questions
func (ctrl *questionHandler) AdminPostOne(c *gin.Context) {
	var req PostQuestionRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.usecase.CreateOne(usecase.CreateOneDto(req), ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, ToQuestionResponse(question))
}

// PUT /api/admin/questions/:question_id
func (ctrl *questionHandler) AdminPutOne(c *gin.Context) {
	var uri QuestionUri
	var req PutQuestionRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	question, err := ctrl.usecase.UpdateOne(usecase.UpdateOneDto{
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
func (ctrl *questionHandler) AdminDeleteOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := ctrl.usecase.DeleteOne(usecase.DeleteOneDto{Id: uri.QuestionId}, ctrl.db); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, nil)
}

// PATCH /api/admin/questions/:question_id
func (ctrl *questionHandler) AdminRestoreOne(c *gin.Context) {
	var uri QuestionUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.usecase.RestoreOne(usecase.RestoreOneDto{Id: uri.QuestionId}, ctrl.db)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}
