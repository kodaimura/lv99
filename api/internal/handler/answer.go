package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"lv99/internal/helper"
	answerModule "lv99/internal/module/answer"
	usecase "lv99/internal/usecase/answer"
)

// -----------------------------
// DTO（Response）
// -----------------------------

type AnswerResponse struct {
	Id         int            `json:"id"`
	QuestionId int            `json:"question_id"`
	AccountId  int            `json:"account_id"`
	CodeDef    string         `json:"code_def"`
	CodeCall   string         `json:"code_call"`
	CallOutput string         `json:"call_output"`
	CallError  string         `json:"call_error"`
	IsCorrect  bool           `json:"is_correct"`
	CorrectAt  *time.Time     `json:"correct_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func ToAnswerResponse(m answerModule.Answer) AnswerResponse {
	return AnswerResponse{
		Id:         m.Id,
		QuestionId: m.QuestionId,
		AccountId:  m.AccountId,
		CodeDef:    m.CodeDef,
		CodeCall:   m.CodeCall,
		CallOutput: m.CallOutput,
		CallError:  m.CallError,
		IsCorrect:  m.IsCorrect,
		CorrectAt:  m.CorrectAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func ToAnswerResponseList(models []answerModule.Answer) []AnswerResponse {
	res := make([]AnswerResponse, 0, len(models))
	for _, m := range models {
		res = append(res, ToAnswerResponse(m))
	}
	return res
}

// -----------------------------
// DTO（Request）
// -----------------------------

type AnswerUri struct {
	AnswerId int `uri:"answer_id" binding:"required"`
}

type QuestionAnswerUri struct {
	QuestionId int `uri:"question_id" binding:"required"`
	AnswerId   int `uri:"answer_id" binding:"required"`
}

type GetAnswersRequest struct {
	QuestionId int `form:"question_id"`
}

type PostAnswersRequest struct {
	QuestionId int    `json:"question_id" binding:"required"`
	CodeDef    string `json:"code_def" binding:"required"`
	CodeCall   string `json:"code_call" binding:"required"`
}

type PutAnswerRequest struct {
	CodeDef  string `json:"code_def" binding:"required"`
	CodeCall string `json:"code_call" binding:"required"`
}

type AdminGetAnswersRequest struct {
	AccountId  int `form:"account_id"`
	QuestionId int `form:"question_id"`
}

// -----------------------------
// Handler Interface
// -----------------------------

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
	usecase usecase.Usecase
}

func NewAnswerHandler(db *gorm.DB, usecase usecase.Usecase) AnswerHandler {
	return &answerHandler{
		usecase: usecase,
	}
}

// -----------------------------
// Handler Implementations
// -----------------------------

// GET /api/answers?question_id=:question_id
func (ctrl *answerHandler) ApiGet(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req GetAnswersRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := ctrl.usecase.Get(usecase.GetDto{
		QuestionId: req.QuestionId,
		AccountId:  accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponseList(answers))
}

// POST /api/answers
func (ctrl *answerHandler) ApiPostOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var req PostAnswersRequest
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.usecase.CreateOne(usecase.CreateOneDto{
		QuestionId: req.QuestionId,
		AccountId:  accountId,
		CodeDef:    req.CodeDef,
		CodeCall:   req.CodeCall,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, ToAnswerResponse(answer))
}

// GET /api/answers/:answer_id
func (ctrl *answerHandler) ApiGetOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.usecase.GetOne(usecase.GetOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponse(answer))
}

// PUT /api/answers/:answer_id
func (ctrl *answerHandler) ApiPutOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	var req PutAnswerRequest
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}
	if err := helper.BindJSON(c, &req); err != nil {
		c.Error(err)
		return
	}

	answer, err := ctrl.usecase.UpdateOne(usecase.UpdateOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
		CodeDef:   req.CodeDef,
		CodeCall:  req.CodeCall,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponse(answer))
}

// DELETE /api/answers/:answer_id
func (ctrl *answerHandler) ApiDeleteOne(c *gin.Context) {
	accountId := helper.GetAccountId(c)
	var uri AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	err := ctrl.usecase.DeleteOne(usecase.DeleteOneDto{
		Id:        uri.AnswerId,
		AccountId: accountId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, nil)
}

// GET /api/admin/answers?account_id=:account_id&question_id=:question_id
func (ctrl *answerHandler) AdminGet(c *gin.Context) {
	var req AdminGetAnswersRequest
	if err := helper.BindQuery(c, &req); err != nil {
		c.Error(err)
		return
	}

	answers, err := ctrl.usecase.Get(usecase.GetDto(req))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponseList(answers))
}

// GET /api/admin/answers/:answer_id
func (ctrl *answerHandler) AdminGetOne(c *gin.Context) {
	var uri AnswerUri
	if err := helper.BindUri(c, &uri); err != nil {
		c.Error(err)
		return
	}

	ans, err := ctrl.usecase.GetOne(usecase.GetOneDto{
		Id: uri.AnswerId,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ToAnswerResponse(ans))
}
