package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetUser godoc
// @Summary Get a user
// @Description Get user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /v2/user [get]
func GetUser(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	}).Info("Request received")

	c.JSON(http.StatusOK, gin.H{
		"version": "v2",
		"user":    "Jane Smith",
	})
}
