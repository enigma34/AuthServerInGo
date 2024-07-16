package middleware

import (
	v1 "AuthServerInGo/handlers/v1"
	"AuthServerInGo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenRevokeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" || services.IsTokenRevoked(token) {
			c.JSON(http.StatusUnauthorized, v1.ErrorResponse{Error: "Token has been revoked."})
			c.Abort()
			return
		}
		c.Next()
	}
}
