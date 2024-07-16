package main

import (
	"AuthServerInGo/routes"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /
func main() {
	logrus.Info("Starting server on :8080")
	InitializeLogger()
	router := gin.Default()

	// fmt.Println(len(services.GetAllUsers()))
	// fmt.Println(services.GetUser("abc@xyz.com").Id)

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// // Version 1 routes
	// v1Routes := router.Group("/v1")
	// {
	// 	v1Routes.GET("/getallusers", v1.GetUser)

	// 	// // @Router /v1/register [post]
	// 	v1Routes.POST("/register", v1.Register)

	// 	// // @Router /v1/login [post]
	// 	v1Routes.POST("/login", v1.Login)
	// }

	// //Version 2 routes
	// v2Routes := router.Group("/v2")
	// {
	// 	v2Routes.GET("/user", v2.GetUser)
	// }

	routes.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		logrus.Fatal("Failed to start server", err)
	}
}
