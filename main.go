package main

import (
	_ "AuthServerInGo/docs"
	v1 "AuthServerInGo/handlers/v1"
	v2 "AuthServerInGo/handlers/v2"
	"AuthServerInGo/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeLogger() {
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    100,
		MaxBackups: 4,
		MaxAge:     1,
		Compress:   true,
	})
}

// @title Auth Server API
// @version 1.0
// @description This is an authentication server API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	InitializeLogger()
	router := gin.Default()

	fmt.Println(len(services.GetAllUsers()))
	fmt.Println(services.GetUser("abc@xyz.com").Id)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Version 1 routes
	v1Routes := router.Group("/v1")
	{
		v1Routes.GET("/user", v1.GetUser)
	}

	// Version 2 routes
	v2Routes := router.Group("/v2")
	{
		v2Routes.GET("/user", v2.GetUser)
	}

	if err := router.Run(":8080"); err != nil {
		logrus.Fatal("Failed to start server", err)
	}
}
