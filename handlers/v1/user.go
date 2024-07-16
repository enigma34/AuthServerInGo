package v1

import (
	"AuthServerInGo/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetUsers retrieves all registered users
// @Summary Retrieves all users
// @Description Returns a JSON list of all users registered in the system.
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} models.User "List of all users"
// @Failure 500 {object} ErrorResponse "Internal Server Error, failed to retrieve users"
// @Security ApiKeyAuth
// @Router /v1/getallusers [get]
func GetUser(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"path":   c.Request.URL.Path,
	}).Info("Request received")

	// c.JSON(http.StatusOK, gin.H{
	// 	"version": "v1",
	// 	"user":    "John Doe",
	// })
	var allUsers = services.GetAllUsers()
	if allUsers == nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to retrieve users"})
		return
	}
	// userJson, err := json.Marshal(allUsers)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to serialize users"})
	// 	return
	// }
	c.JSON(http.StatusOK, allUsers)
}
