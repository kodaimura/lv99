package router

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/domain/account"
	"lv99/internal/domain/account_profile"
	"lv99/internal/domain/answer"
	"lv99/internal/domain/auth"
	"lv99/internal/domain/chat"
	"lv99/internal/domain/comment"
	"lv99/internal/domain/executor"
	"lv99/internal/domain/question"
	"lv99/internal/infrastructure/db"
	"lv99/internal/infrastructure/externalapi"
)

var gorm = db.NewGormDB()
var sqlx = db.NewSqlxDB()

/* DI (Repository) */
var accountRepository = account.NewRepository()
var accountProfileRepository = account_profile.NewRepository()
var questionRepository = question.NewRepository()
var answerRepository = answer.NewRepository()
var commentRepository = comment.NewRepository()
var chatRepository = chat.NewRepository()

/* DI (Query) */
var accountQuery = account.NewQuery(sqlx)
var chatQuery = chat.NewQuery(sqlx)

/* DI (Service) */
var authService = auth.NewService(accountRepository, accountProfileRepository)
var executorService = executor.NewService(externalapi.NewHttpCodeExecutor())
var accountService = account.NewService(accountRepository, accountQuery)
var accountProfileService = account_profile.NewService(accountProfileRepository)
var questionService = question.NewService(questionRepository)
var answerService = answer.NewService(answerRepository, questionService, executorService)
var commentService = comment.NewService(commentRepository)
var chatService = chat.NewService(chatRepository, chatQuery)

/* DI (Controller) */
var authController = auth.NewController(gorm, authService, accountProfileService)
var accountController = account.NewController(gorm, accountService)
var accountProfileController = account_profile.NewController(gorm, accountProfileService)
var questionController = question.NewController(gorm, questionService)
var answerController = answer.NewController(gorm, answerService)
var commentController = comment.NewController(gorm, commentService)
var chatController = chat.NewController(gorm, chatService)

func SetApi(r *gin.RouterGroup) {
	r.Use(ApiErrorHandler())
	r.POST("/accounts/signup", authController.ApiSignup)
	r.POST("/accounts/login", authController.ApiLogin)
	r.POST("/accounts/refresh", authController.ApiRefresh)
	r.POST("/accounts/logout", authController.ApiLogout)

	auth := r.Group("", ApiAuthMiddleware())
	{
		auth.PUT("/accounts/me/password", authController.ApiPutMePassword)
		auth.GET("/accounts/me", accountController.ApiGetMe)
		auth.PUT("/accounts/me", accountController.ApiPutMe)
		auth.DELETE("/accounts/me", accountController.ApiDeleteMe)

		auth.GET("/accounts/me/profile", accountProfileController.ApiGetMe)
		auth.PUT("/accounts/me/profile", accountProfileController.ApiPutMe)

		auth.GET("/questions", questionController.ApiGet)
		auth.GET("/questions/:question_id", questionController.ApiGetOne)

		auth.GET("/questions/:question_id/answers", answerController.ApiGet)
		auth.POST("/questions/:question_id/answers", answerController.ApiPostOne)
		auth.GET("/answers/:answer_id", answerController.ApiGetOne)
		auth.PUT("/answers/:answer_id", answerController.ApiPutOne)
		auth.DELETE("/answers/:answer_id", answerController.ApiDeleteOne)

		auth.GET("/answers/:answer_id/comments", commentController.ApiGet)
		auth.POST("/answers/:answer_id/comments", commentController.ApiPostOne)
		auth.GET("/comments/:comment_id", commentController.ApiGetOne)
		auth.PUT("/comments/:comment_id", commentController.ApiPutOne)
		auth.DELETE("/comments/:comment_id", commentController.ApiDeleteOne)

		auth.GET("/chats/ws", chatController.WsConnect)
		auth.GET("/chats/:to_id", chatController.ApiGet)
	}

	admin := r.Group("admin", ApiAuthMiddleware())
	{
		admin.GET("/accounts/with-profile", accountController.AdminGetWithProfile)
		admin.GET("/accounts/:account_id/with-profile", accountController.AdminGetOneWithProfile)

		admin.GET("/questions", questionController.AdminGet)
		admin.POST("/questions", questionController.AdminPostOne)
		admin.GET("/questions/:question_id", questionController.AdminGetOne)
		admin.PUT("/questions/:question_id", questionController.AdminPutOne)
		admin.DELETE("/questions/:question_id", questionController.AdminDeleteOne)
		admin.PATCH("/questions/:question_id", questionController.AdminRestoreOne)
	}
}
