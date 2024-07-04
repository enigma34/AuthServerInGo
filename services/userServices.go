package services

import (
	"AuthServerInGo/models"
)

var users = []models.User{
	{Id: 1, Email: "abc@xyz.com", PasswordHash: "$2a$11$BDW8s0ctkCo35NqKfdXmG.aSME0Tqne6wepyeVYctpkeft2KStluC"},
	{Id: 2, Email: "123@xyz.com", PasswordHash: "$2a$11$BDW8s0ctkCo35NqKfdXmG.aSME0Tqne6wepyeVYctpkeft2KStluC"},
	{Id: 3, Email: "wsad@xyz.com", PasswordHash: "$2a$11$BDW8s0ctkCo35NqKfdXmG.aSME0Tqne6wepyeVYctpkeft2KStluC"},
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
