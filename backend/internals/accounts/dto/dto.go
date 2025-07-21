package dto

type LoginStruct struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role string `json:"role" binding:"required"`
}

type UserAccount struct {
	Email    string `json:"email" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

type RoleStruct struct {
	Name    string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}