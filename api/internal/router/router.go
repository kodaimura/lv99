package router

import (
	"github.com/gin-gonic/gin"

	feature_account "lv99/internal/feature/account"
	feature_answer "lv99/internal/feature/answer"
	feature_chat "lv99/internal/feature/chat"
	feature_comment "lv99/internal/feature/comment"
	"lv99/internal/handler"
	"lv99/internal/infrastructure/db"
	"lv99/internal/infrastructure/externalapi"
	"lv99/internal/module/account"
	"lv99/internal/module/account_profile"
	"lv99/internal/module/answer"
	"lv99/internal/module/auth"
	"lv99/internal/module/chat"
	"lv99/internal/module/comment"
	"lv99/internal/module/executor"
	"lv99/internal/module/question"
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
var chatQuery = chat.NewQuery(sqlx)

var featureAccountQuery = feature_account.NewQuery(sqlx)
var featureAnswerQuery = feature_answer.NewQuery(sqlx)
var featureCommentQuery = feature_comment.NewQuery(sqlx)
var featureChatQuery = feature_chat.NewQuery(sqlx)

/* DI (Service) */
var authService = auth.NewService(accountRepository, accountProfileRepository)
var executorService = executor.NewService(externalapi.NewHttpCodeExecutor())
var accountService = account.NewService(accountRepository)
var accountProfileService = account_profile.NewService(accountProfileRepository)
var questionService = question.NewService(questionRepository)
var answerService = answer.NewService(answerRepository, questionService, executorService)
var commentService = comment.NewService(commentRepository)
var chatService = chat.NewService(chatRepository, chatQuery)

var featureAccountService = feature_account.NewService(featureAccountQuery)
var featureAnswerService = feature_answer.NewService(featureAnswerQuery)
var featureCommentService = feature_comment.NewService(featureCommentQuery)
var featureChatService = feature_chat.NewService(featureChatQuery)

/* DI (Controller) */
var authController = handler.NewAuthHandler(gorm, authService, accountProfileService)
var accountController = handler.NewAccountHandler(gorm, accountService)
var accountProfileController = handler.NewAccountProfileHandler(gorm, accountProfileService)
var questionController = handler.NewQuestionHandler(gorm, questionService)
var answerController = handler.NewAnswerHandler(gorm, answerService)
var commentController = handler.NewCommentHandler(gorm, commentService)
var chatController = handler.NewChatHandler(gorm, chatService)

var featureAccountController = feature_account.NewController(gorm, featureAccountService)
var featureAnswerController = feature_answer.NewController(gorm, featureAnswerService)
var featureCommentController = feature_comment.NewController(gorm, featureCommentService)
var featureChatController = feature_chat.NewController(gorm, featureChatService)

func SetApi(r *gin.RouterGroup) {
	r.Use(ApiErrorHandler())
	//r.POST("/accounts/signup", authController.ApiSignup)
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

		auth.GET("/answers", answerController.ApiGet)
		auth.POST("/answers", answerController.ApiPostOne)
		auth.GET("/answers/:answer_id", answerController.ApiGetOne)
		auth.PUT("/answers/:answer_id", answerController.ApiPutOne)
		auth.DELETE("/answers/:answer_id", answerController.ApiDeleteOne)
		auth.GET("/answers/status", featureAnswerController.ApiGetStatus)

		auth.GET("/comments", commentController.ApiGet)
		auth.POST("/comments", commentController.ApiPostOne)
		auth.GET("/comments/:comment_id", commentController.ApiGetOne)
		auth.PUT("/comments/:comment_id", commentController.ApiPutOne)
		auth.DELETE("/comments/:comment_id", commentController.ApiDeleteOne)
		auth.GET("/comments/with-profile", featureCommentController.ApiGetWithProfile)
		auth.GET("/comments/count", featureCommentController.ApiGetRecentCount)

		auth.GET("/chats/unread-count", featureChatController.ApiGetUnreadCount)
		auth.GET("/chats/:to_id", chatController.ApiGet)
		auth.PUT("/chats/read", chatController.ApiRead)

		auth.GET("/accounts/admin/with-profile", featureAccountController.ApiGetAdminWithProfile)
	}

	admin := r.Group("admin", ApiAuthMiddleware())
	{
		admin.POST("/accounts/signup", authController.ApiSignup)
		admin.GET("/accounts/with-profile", featureAccountController.AdminGetWithProfile)
		admin.GET("/accounts/:account_id/with-profile", featureAccountController.AdminGetOneWithProfile)

		admin.GET("/questions", questionController.AdminGet)
		admin.POST("/questions", questionController.AdminPostOne)
		admin.GET("/questions/:question_id", questionController.AdminGetOne)
		admin.PUT("/questions/:question_id", questionController.AdminPutOne)
		admin.DELETE("/questions/:question_id", questionController.AdminDeleteOne)
		admin.PATCH("/questions/:question_id", questionController.AdminRestoreOne)

		admin.GET("/answers", answerController.AdminGet)
		admin.GET("/answers/:answer_id", answerController.AdminGetOne)
		admin.GET("/answers/search", featureAnswerController.AdminSearch)
		admin.GET("/answers/status", featureAnswerController.AdminGetStatus)

		admin.GET("/comments/count", featureCommentController.AdminGetRecentCount)
	}
}

func SetWs(r *gin.RouterGroup) {
	auth := r.Group("", ApiAuthMiddleware())
	{
		auth.GET("/chats", chatController.WsConnect)
	}
}
