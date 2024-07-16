package dtos

type RegisterRequest struct {
	Email    string   `json:"email" binding:"required,email"`
	Password string   `json:"password" binding:"required,max=16"`
	Roles    []string `json:"roles"`
}
