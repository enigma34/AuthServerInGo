package models

type User struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
