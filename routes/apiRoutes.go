package routes

import (
	_ "AuthServerInGo/docs"
	v1 "AuthServerInGo/handlers/v1"
	v2 "AuthServerInGo/handlers/v2"
	"AuthServerInGo/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Version 1 routes
	v1Routes := r.Group("/v1")
	{
		v1Routes.GET("/getallusers", middleware.AuthMiddleware(), middleware.TokenRevokeMiddleware(), v1.GetUser)

		// // @Router /v1/register [post]
		v1Routes.POST("/register", v1.Register)

		// // @Router /v1/login [post]
		v1Routes.POST("/login", v1.Login)

		v1Routes.POST("/revoke", v1.RevokeToken)
	}

	//Version 2 routes
	v2Routes := r.Group("/v2")
	{
		v2Routes.GET("/user", v2.GetUser)
	}

}
