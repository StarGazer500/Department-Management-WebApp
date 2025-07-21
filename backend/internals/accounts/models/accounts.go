package models

// import (
// 	"database/sql"
	
// )

type Users struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type UserRole struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserWithRoles struct {
	User  Users  `json:"user"`
	Roles []Role `json:"roles"`
}

// type AccountModel struct{
// 	db *sql.DB 
// }




