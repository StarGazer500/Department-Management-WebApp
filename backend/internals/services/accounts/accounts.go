package accounts

import (
	
	// "errors"
	"fmt"
    "github.com/gin-gonic/gin"
	"github.com/StarGazer500/Department-Management-WebApp/internals/database"
	"github.com/StarGazer500/Department-Management-WebApp/internals/repository/accounts"
	"github.com/StarGazer500/Department-Management-WebApp/internals/utils"
)

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


func LoginService(login *LoginStruct, ctx *gin.Context) (string) {
	data, err := accounts.GetUserWithRoles(database.Dbinstance, login.Email)
	if err != nil {
		fmt.Println(err)
		return "Username not found"
	}

	claim:=utils.TokenClaimStruct{
		MyAuthServer:    "AuthServer",
		AuthUserEmail:   data.User.Email,
		AuthUserId:      data.User.ID,

	} 

	if claim.MyAuthServer=="" || claim.AuthUserEmail =="" || claim.AuthUserId==0{
		fmt.Println("Token Struct is invalid",claim)
		return "something bad happened"
	}
	for _, item := range data.Roles {
		fmt.Println(item.Name)
		if item.Name == login.Role && utils.ComparePassword(login.Password, data.User.PasswordHash) {
			
			if item.Name == "lecturer" {
				fmt.Println("lecturer Credential is working")
				access_token, error1 := utils.GenerateAccessToken(claim)
				refresh_token, error2 := utils.GenerateRefreshToken(claim)
				if error1 != nil || error2 != nil {
					fmt.Println("token generation error")
					return "something bad happened"
				}
				ctx.SetCookie("access", access_token, 3600, "/", "127.0.0.1", false, true)
				ctx.SetCookie("refresh", refresh_token, 3600, "/", "127.0.0.1", false, true)



				
				return "user is lecturer"
			} else if item.Name == "admin" {
				fmt.Println("Admin Credential is working")
				access_token, error1 := utils.GenerateAccessToken(claim)
				refresh_token, error2 := utils.GenerateRefreshToken(claim)
				if error1 != nil || error2 != nil {
					fmt.Println("token generation error")
					return "something bad happened"
				}
				ctx.SetCookie("access", access_token, 3600, "/", "127.0.0.1", false, true)
				ctx.SetCookie("refresh", refresh_token, 3600, "/", "127.0.0.1", false, true)
				return "user is admin"
			} else if item.Name == "technicians" {
				fmt.Println("Technician Credential is working")
				access_token, error1 := utils.GenerateAccessToken(claim)
				refresh_token, error2 := utils.GenerateRefreshToken(claim)
				if error1 != nil || error2 != nil {
					fmt.Println("token generation error")
					return "something bad happened"
				}
				ctx.SetCookie("access", access_token, 3600, "/", "127.0.0.1", false, true)
				ctx.SetCookie("refresh", refresh_token, 3600, "/", "127.0.0.1", false, true)
				return "user is technicians"
			} else if item.Name == "student"  {
				fmt.Println("student Credential is working")
				access_token, error1 := utils.GenerateAccessToken(claim)
				refresh_token, error2 := utils.GenerateRefreshToken(claim)
				if error1 != nil || error2 != nil {
					fmt.Println("token generation error")
					return "something bad happened"
				}
				ctx.SetCookie("access", access_token, 3600, "/", "127.0.0.1", false, true)
				ctx.SetCookie("refresh", refresh_token, 3600, "/", "127.0.0.1", false, true)
				return "user is student"
			} else if item.Name == "teaching assistant"  {
				fmt.Println("teaching assistant Credential is working")
				access_token, error1 := utils.GenerateAccessToken(claim)
				refresh_token, error2 := utils.GenerateRefreshToken(claim)
				if error1 != nil || error2 != nil {
					fmt.Println("token generation error")
					return "something bad happened"
				}
				ctx.SetCookie("access", access_token, 3600, "/", "127.0.0.1", false, true)
				ctx.SetCookie("refresh", refresh_token, 3600, "/", "127.0.0.1", false, true)
				return "user is teaching"
			} else if item.Name == "hod" {
				fmt.Println("hod Credential is working")
				access_token, error1 := utils.GenerateAccessToken(claim)
				refresh_token, error2 := utils.GenerateRefreshToken(claim)
				if error1 != nil || error2 != nil {
					fmt.Println("token generation error")
					return "something bad happened"
				}
				ctx.SetCookie("access", access_token, 3600, "/", "127.0.0.1", false, true)
				ctx.SetCookie("refresh", refresh_token, 3600, "/", "127.0.0.1", false, true)
				return "user is hod"
			} else if item.Name == "secretary"{
				fmt.Println("secretary Credential is working")
				access_token, error1 := utils.GenerateAccessToken(claim)
				refresh_token, error2 := utils.GenerateRefreshToken(claim)
				if error1 != nil || error2 != nil {
					fmt.Println("token generation error")
					return "something bad happened"
				}
				ctx.SetCookie("access", access_token, 3600, "/", "127.0.0.1", false, true)
				ctx.SetCookie("refresh", refresh_token, 3600, "/", "127.0.0.1", false, true)
				return "user is secretary"
			} else if item.Name == "Instrument Officer" {
				fmt.Println("Instrument Officer Credential is working")
				access_token, error1 := utils.GenerateAccessToken(claim)
				refresh_token, error2 := utils.GenerateRefreshToken(claim)
				if error1 != nil || error2 != nil {
					fmt.Println("token generation error")
					return "something bad happened"
				}
				ctx.SetCookie("access", access_token, 3600, "/", "127.0.0.1", false, true)
				ctx.SetCookie("refresh", refresh_token, 3600, "/", "127.0.0.1", false, true)
				return "user is instrument officer"
			}
		}
	}
	return "Invalid User"
}

func CreateUserService(userAccount * UserAccount)(string){
	hashedPassword,errorMessage:= utils.HashPassword(userAccount.PasswordHash)
	if errorMessage!=nil{
		return"User not Created"
	}
	_, error := accounts.CreateUser(database.Dbinstance, userAccount.Email,hashedPassword,userAccount.FirstName,userAccount.LastName,userAccount.RoleName)
		if error != nil {
			fmt.Println(error)
			
			return "User not created"
		}
	return "User Created Successfully"

}

func CreateRoleService(role * RoleStruct)(string){
	_, error := accounts.CreateRole(database.Dbinstance, role.Name,role.Description)
		if error != nil {
			
			return "Role not created"
		}
	return "Role Created Successfully"

}