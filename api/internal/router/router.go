package router

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/adapter/db"

	"lv99/internal/module/account"
	"lv99/internal/module/account_profile"
	"lv99/internal/module/answer"
	"lv99/internal/module/chat"
	"lv99/internal/module/comment"
	"lv99/internal/module/question"

	account_uc "lv99/internal/usecase/account"
	account_extended_uc "lv99/internal/usecase/account_extended"
	account_profile_uc "lv99/internal/usecase/account_profile"
	answer_uc "lv99/internal/usecase/answer"
	answer_extended_uc "lv99/internal/usecase/answer_extended"
	auth_uc "lv99/internal/usecase/auth"
	chat_uc "lv99/internal/usecase/chat"
	chat_extended_uc "lv99/internal/usecase/chat_extended"
	comment_uc "lv99/internal/usecase/comment"
	comment_extended_uc "lv99/internal/usecase/comment_extended"
	question_uc "lv99/internal/usecase/question"

	account_h "lv99/internal/handler/account"
	account_extended_h "lv99/internal/handler/account_extended"
	account_profile_h "lv99/internal/handler/account_profile"
	answer_h "lv99/internal/handler/answer"
	answer_extended_h "lv99/internal/handler/answer_extended"
	auth_h "lv99/internal/handler/auth"
	chat_h "lv99/internal/handler/chat"
	chat_extended_h "lv99/internal/handler/chat_extended"
	comment_h "lv99/internal/handler/comment"
	comment_extended_h "lv99/internal/handler/comment_extended"
	question_h "lv99/internal/handler/question"
)

var gorm = db.NewGormDB()
var sqlx = db.NewSqlxDB()

/* DI (Repository) */
var accountRepository = account.NewRepository()
var accountProfileRepository = account_profile.NewRepository()
var answerRepository = answer.NewRepository()
var chatRepository = chat.NewRepository()
var commentRepository = comment.NewRepository()
var questionRepository = question.NewRepository()

/* DI (Service) */
var accountService = account.NewService(accountRepository)
var accountProfileService = account_profile.NewService(accountProfileRepository)
var answerService = answer.NewService(answerRepository)
var chatService = chat.NewService(chatRepository)
var commentService = comment.NewService(commentRepository)
var questionService = question.NewService(questionRepository)

/* DI (Usecase) */
var authUsecase = auth_uc.NewUsecase(gorm, accountService, accountProfileService)
var accountUsecase = account_uc.NewUsecase(gorm, accountService)
var accountProfileUsecase = account_profile_uc.NewUsecase(gorm, accountProfileService)
var questionUsecase = question_uc.NewUsecase(gorm, questionService)
var answerUsecase = answer_uc.NewUsecase(gorm, answerService, questionService)
var chatUsecase = chat_uc.NewUsecase(gorm, sqlx, chatService)
var commentUsecase = comment_uc.NewUsecase(gorm, commentService)

var accountExUsecase = account_extended_uc.NewUsecase(gorm, sqlx)
var answerExUsecase = answer_extended_uc.NewUsecase(gorm, sqlx)
var chatExUsecase = chat_extended_uc.NewUsecase(gorm, sqlx)
var commentExUsecase = comment_extended_uc.NewUsecase(gorm, sqlx)

/* DI (Handler) */
var accountHandler = account_h.NewHandler(accountUsecase)
var accountProfileHandler = account_profile_h.NewHandler(accountProfileUsecase)
var authHandler = auth_h.NewHandler(authUsecase)
var answerHandler = answer_h.NewHandler(answerUsecase)
var chatHandler = chat_h.NewHandler(chatUsecase)
var commentHandler = comment_h.NewHandler(commentUsecase)
var questionHandler = question_h.NewHandler(questionUsecase)

var accountExHandler = account_extended_h.NewHandler(accountExUsecase)
var answerExHandler = answer_extended_h.NewHandler(answerExUsecase)
var chatExHandler = chat_extended_h.NewHandler(chatExUsecase)
var commentExHandler = comment_extended_h.NewHandler(commentExUsecase)

func SetApi(r *gin.RouterGroup) {
	r.Use(ApiErrorHandler())
	//r.POST("/accounts/signup", authHandler.ApiSignup)
	r.POST("/accounts/login", authHandler.ApiLogin)
	r.POST("/accounts/refresh", authHandler.ApiRefresh)
	r.POST("/accounts/logout", authHandler.ApiLogout)

	auth := r.Group("", ApiAuthMiddleware())
	{
		auth.PUT("/accounts/me/password", authHandler.ApiPutMePassword)
		auth.GET("/accounts/me", accountHandler.ApiGetMe)
		auth.PUT("/accounts/me", accountHandler.ApiPutMe)
		auth.DELETE("/accounts/me", accountHandler.ApiDeleteMe)

		auth.GET("/accounts/me/profile", accountProfileHandler.ApiGetMe)
		auth.PUT("/accounts/me/profile", accountProfileHandler.ApiPutMe)

		auth.GET("/questions", questionHandler.ApiGet)
		auth.GET("/questions/:question_id", questionHandler.ApiGetOne)

		auth.GET("/answers", answerHandler.ApiGet)
		auth.POST("/answers", answerHandler.ApiPostOne)
		auth.GET("/answers/:answer_id", answerHandler.ApiGetOne)
		auth.PUT("/answers/:answer_id", answerHandler.ApiPutOne)
		auth.DELETE("/answers/:answer_id", answerHandler.ApiDeleteOne)
		auth.GET("/answers/status", answerExHandler.ApiGetStatus)

		auth.GET("/comments", commentHandler.ApiGet)
		auth.POST("/comments", commentHandler.ApiPostOne)
		auth.GET("/comments/:comment_id", commentHandler.ApiGetOne)
		auth.PUT("/comments/:comment_id", commentHandler.ApiPutOne)
		auth.DELETE("/comments/:comment_id", commentHandler.ApiDeleteOne)
		auth.GET("/comments/with-profile", commentExHandler.ApiGetWithProfile)
		auth.GET("/comments/count", commentExHandler.ApiGetRecentCount)

		auth.GET("/chats/unread-count", chatExHandler.ApiGetUnreadCount)
		auth.GET("/chats/:to_id", chatHandler.ApiGet)
		auth.PUT("/chats/read", chatHandler.ApiRead)

		auth.GET("/accounts/admin/with-profile", accountExHandler.ApiGetAdminWithProfile)
	}

	admin := r.Group("/admin", ApiAuthMiddleware())
	{
		admin.POST("/accounts/signup", authHandler.ApiSignup)
		admin.GET("/accounts/with-profile", accountExHandler.AdminGetWithProfile)
		admin.GET("/accounts/:account_id/with-profile", accountExHandler.AdminGetOneWithProfile)

		admin.GET("/questions", questionHandler.AdminGet)
		admin.POST("/questions", questionHandler.AdminPostOne)
		admin.GET("/questions/:question_id", questionHandler.AdminGetOne)
		admin.PUT("/questions/:question_id", questionHandler.AdminPutOne)
		admin.DELETE("/questions/:question_id", questionHandler.AdminDeleteOne)
		admin.PATCH("/questions/:question_id", questionHandler.AdminRestoreOne)

		admin.GET("/answers", answerHandler.AdminGet)
		admin.GET("/answers/:answer_id", answerHandler.AdminGetOne)
		admin.GET("/answers/search", answerExHandler.AdminSearch)
		admin.GET("/answers/status", answerExHandler.AdminGetStatus)

		admin.GET("/comments/count", commentExHandler.AdminGetRecentCount)
	}
}

func SetWs(r *gin.RouterGroup) {
	auth := r.Group("", ApiAuthMiddleware())
	{
		auth.GET("/chats", chatHandler.WsConnect)
	}
}
