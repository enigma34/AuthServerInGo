package services

import (
	"AuthServerInGo/internal/config"
	"AuthServerInGo/models"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var users = []models.User{
	{Id: 1, Email: "abc@xyz.com", PasswordHash: "$2a$11$BDW8s0ctkCo35NqKfdXmG.aSME0Tqne6wepyeVYctpkeft2KStluC", CreatedAt: time.Now(), Roles: []string{"admin", "user"}},
	{Id: 2, Email: "123@xyz.com", PasswordHash: "$2a$11$BDW8s0ctkCo35NqKfdXmG.aSME0Tqne6wepyeVYctpkeft2KStluC", CreatedAt: time.Now(), Roles: []string{"user"}},
	{Id: 3, Email: "wsad@xyz.com", PasswordHash: "$2a$11$BDW8s0ctkCo35NqKfdXmG.aSME0Tqne6wepyeVYctpkeft2KStluC", CreatedAt: time.Now(), Roles: []string{"admin"}},
}

var _revokedTokens = []string{}

type CustomClaims struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func GetAllUsers() []models.User {
	return users
}

func GetUser(email string) *models.User {
	for _, user := range users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

func AddUser(user models.User) {
	users = append(users, user)
}

func CheckUserInDb(email string) bool {
	var allUsers = GetAllUsers()
	//checkuserexists bool =
	for _, checkuserexists := range allUsers {
		if checkuserexists.Email == email {
			return true
		}
	}
	return false
}

func GetUsersCount() int {
	return len(GetAllUsers())
}

func CreateToken(user models.User) (string, error) {

	cfg, err := config.LoadConfig()
	if err != nil {
		//fmt.Println("Error loading config:", err)
		return "Error loading config", err
	}
	var signingKey = []byte(cfg.JWT.SigningKey)
	claims := CustomClaims{
		Email: user.Email,
		Roles: user.Roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(signingKey)
	if err != nil {
		return "Error signing token", err
	}
	return accessToken, nil
}

func IsValidToken(signedToken string) bool {
	cfg, err := config.LoadConfig()
	if err != nil {
		return false
	}
	var signingKey = []byte(cfg.JWT.SigningKey)

	parts := strings.Split(signedToken, " ")
	signedToken = parts[1]

	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return false
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		err = errors.New("couldn't aprse claims")
		return false
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expierd")
		return false
	}
	return true
}

func RevokeToken(signedToken string) {
	_revokedTokens = append(_revokedTokens, signedToken)
}

func IsTokenRevoked(signedToken string) bool {
	parts := strings.Split(signedToken, " ")
	signedToken = parts[1]

	for _, checkTokenExistsInRevokeList := range _revokedTokens {
		if checkTokenExistsInRevokeList == signedToken {
			return true
		}
	}
	return false
}
