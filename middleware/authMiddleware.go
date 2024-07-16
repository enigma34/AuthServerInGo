package middleware

import (
	v1 "AuthServerInGo/handlers/v1"
	"AuthServerInGo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" || !services.IsValidToken(token) {
			c.JSON(http.StatusUnauthorized, v1.ErrorResponse{Error: "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
