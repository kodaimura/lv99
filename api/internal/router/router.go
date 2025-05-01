package router

import (
	"github.com/gin-gonic/gin"

	"lv99/internal/controller"
	"lv99/internal/infrastructure/db"
	"lv99/internal/middleware"
	repository "lv99/internal/repository/impl"
	"lv99/internal/service"
)

var gorm = db.NewGormDB()
//var sqlx = db.NewSqlxDB()

/* DI (Repository) */
var accountRepository = repository.NewGormAccountRepository(gorm)
var adminRepository = repository.NewGormAdminRepository(gorm)

/* DI (Query) */
//var xxxQuery = query.NewXxxQuery(sqlx)

/* DI (Service) */
var accountService = service.NewAccountService(accountRepository)
var adminService = service.NewAdminService(adminRepository)

/* DI (Controller) */
var accountController = controller.NewAccountController(accountService)
var adminController = controller.NewAdminController(adminService)

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
	}
}

func SetAdminApi(r *gin.RouterGroup) {
	r.Use(middleware.ApiErrorHandler())

	r.POST("/admins/signup", adminController.ApiSignup)
	r.POST("/admins/login", adminController.ApiLogin)
	r.POST("/admins/refresh", adminController.ApiRefresh)
	r.POST("/admins/logout", adminController.ApiLogout)

	auth := r.Group("", middleware.ApiAdminAuth())
	{
		auth.GET("/admins/me", adminController.ApiGetOne)
		auth.PUT("/admins/me", adminController.ApiPutOne)
		auth.PUT("/admins/me/password", adminController.ApiPutPassword)
		auth.DELETE("/admins/me", adminController.ApiDeleteOne)
	}
}
