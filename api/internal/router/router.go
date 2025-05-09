package router

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/controller"
	"lv99/internal/infrastructure/db"
	"lv99/internal/infrastructure/externalapi"
	"lv99/internal/middleware"
	repository "lv99/internal/repository/impl"
	"lv99/internal/service"
)

var gorm = db.NewGormDB()
//var sqlx = db.NewSqlxDB()

/* DI (Repository) */
var accountRepository = repository.NewGormAccountRepository(gorm)
var questionRepository = repository.NewGormQuestionRepository(gorm)
var answerRepository = repository.NewGormAnswerRepository(gorm)
var commentRepository = repository.NewGormCommentRepository(gorm)

/* DI (Query) */
//var xxxQuery = query.NewXxxQuery(sqlx)

/* DI (Service) */
var accountService = service.NewAccountService(accountRepository)
var questionService = service.NewQuestionService(questionRepository)
var answerService = service.NewAnswerService(questionRepository, answerRepository, externalapi.NewHttpCodeExecutor())
var commentService = service.NewCommentService(commentRepository)

/* DI (Controller) */
var accountController = controller.NewAccountController(accountService)
var questionController = controller.NewQuestionController(questionService)
var answerController = controller.NewAnswerController(answerService)
var commentController = controller.NewCommentController(commentService)


func SetApi(r *gin.RouterGroup) {
	r.Use(middleware.ApiErrorHandler())
	r.POST("/accounts/signup", accountController.ApiSignup)
	r.POST("/accounts/login", accountController.ApiLogin)
	r.POST("/accounts/refresh", accountController.ApiRefresh)
	r.POST("/accounts/logout", accountController.ApiLogout)

	auth := r.Group("", middleware.ApiAuth())
	{
		auth.GET("/accounts/me", accountController.ApiGetOne)
		auth.PUT("/accounts/me", accountController.ApiPutOne)
		auth.PUT("/accounts/me/password", accountController.ApiPutPassword)
		auth.DELETE("/accounts/me", accountController.ApiDeleteOne)

		auth.GET("/questions", questionController.ApiGet)
		auth.GET("/questions/:question_id", questionController.ApiGetOne)

		auth.GET("/questions/:question_id/answers", answerController.ApiGet)
		auth.POST("/questions/:question_id/answers", answerController.ApiPostOne)
		auth.PUT("/questions/:question_id/answers/:answer_id", answerController.ApiPutOne)
		auth.DELETE("/questions/:question_id/answers/:answer_id", answerController.ApiDeleteOne)
		auth.GET("/answers/:answer_id/comments", commentController.ApiGet)
		auth.POST("/answers/:answer_id/comments", commentController.ApiPostOne)
		auth.PUT("/answers/:answer_id/comments/:comment_id", commentController.ApiPutOne)
		auth.DELETE("/answers/:answer_id/comments/:comment_id", commentController.ApiDeleteOne)
	}

	admin := r.Group("admin", middleware.ApiAuth())
	{
		admin.POST("/questions", questionController.AdminPostOne)
		admin.GET("/questions", questionController.AdminGet)
		admin.GET("/questions/:question_id", questionController.AdminGetOne)
		admin.PUT("/questions/:question_id", questionController.AdminPutOne)
		admin.DELETE("/questions/:question_id", questionController.AdminDeleteOne)
		admin.PATCH("/questions/:question_id", questionController.AdminRestoreOne)
	}
}