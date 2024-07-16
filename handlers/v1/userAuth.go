package v1

import (
	"AuthServerInGo/dtos"
	"AuthServerInGo/models"
	"AuthServerInGo/services"
	"AuthServerInGo/utility"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var validRoles = map[string]bool{
	"user":  true,
	"admin": true,
}

// Register a new user
// @Summary Register a new user
// @Description Registers a new user with an email and password, checks if the user exists, hashes the password, and stores the user in the database.
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body dtos.RegisterRequest true "Register Request"
// @Success 200 {object} SuccessResponse "Successfully registered"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Router /v1/register [post]
func Register(c *gin.Context) {
	var registerRequest dtos.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if services.CheckUserInDb(registerRequest.Email) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Email id already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: fmt.Sprintf("Error hashing password: %v", err)})
		return
	}
	for _, role := range registerRequest.Roles {
		if _, exists := validRoles[role]; !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid role provided: %s", role)})
			return
		}
	}
	var user models.User
	user.Id = services.GetUsersCount() + 1
	user.Email = registerRequest.Email
	user.PasswordHash = string(passwordHash)
	user.CreatedAt = time.Now()
	if len(registerRequest.Roles) == 0 {
		user.Roles = []string{"user"}
	} else {
		user.Roles = registerRequest.Roles
	}
	services.AddUser(user)
	c.JSON(http.StatusOK, SuccessResponse{Message: fmt.Sprintf("User: %v registered successfully!!!", user.Id)})
}

// Login authenticates a user
// @Summary User login
// @Description Authenticates a user by email and password, returning an access token if successful.
// @Tags authentication
// @Accept json
// @Produce json
// @Param body body dtos.LoginRequest true "Login Request"
// @Success 200 {object} dtos.LoginResponse "Returns an access token on successful authentication."
// @Failure 400 {object} ErrorResponse "Returns an error if the login request fails due to bad input, user not found, or incorrect password."
// @Router /v1/login [post]
func Login(c *gin.Context) {
	var loginRequest dtos.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	user := services.GetUser(loginRequest.Email)
	if user == nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "User not found"})
		return
	}

	if utility.CheckPasswordHash(loginRequest.Password, user.PasswordHash) == false {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Wrong password"})
		return
	}
	accessToken, err := services.CreateToken(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	var response = dtos.LoginResponse{AccessToken: accessToken}
	c.JSON(http.StatusOK, response)
}

// RevokeToken revokes an access token
// @Summary Revoke access token
// @Description Revokes the access token provided in the request body.
// @Tags token
// @Accept  json
// @Produce  json
// @Param   access_token   body   dtos.RevokeTokenRequest  true  "Access Token"
// @Success 200  {object}  SuccessResponse  "Token revoked successfully"
// @Failure 400  {object}  ErrorResponse  "Bad Request"
// @Failure 401  {object}  ErrorResponse  "Unauthorized"
// @Failure 500  {object}  ErrorResponse  "Internal Server Error"
// @Router /v1/revoke [post]
func RevokeToken(c *gin.Context) {
	var revokeToken dtos.RevokeTokenRequest
	if err := c.ShouldBindJSON(&revokeToken); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request"})
		return
	}
	services.RevokeToken(revokeToken.AccessToken)
	c.JSON(http.StatusOK, SuccessResponse{Message: "Token revoked"})
}
