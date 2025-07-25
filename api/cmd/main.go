package main

import (
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"lv99/config"
	"lv99/internal/adapter/auth"
	"lv99/internal/adapter/externalapi"
	"lv99/internal/adapter/file"
	"lv99/internal/adapter/logger"
	"lv99/internal/adapter/mailer"
	"lv99/internal/core"
	"lv99/internal/router"
)

func main() {
	f1 := file.GetAccessLogFile()
	f2 := file.GetAppLogFile()
	defer f1.Close()
	defer f2.Close()

	gin.DefaultWriter = io.MultiWriter(os.Stdout, f1)

	core.SetLogger(logger.NewMultiLogger(f2))
	core.SetMailer(mailer.NewSmtpMailer())
	core.SetAuth(auth.NewJwtAuth())
	core.SetCodeExecutor(externalapi.NewHttpCodeExecutor())

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.FrontendOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.SetApi(r.Group("/api"))
	router.SetWs(r.Group("/ws"))
	r.Run(":" + config.AppPort)
}
