package accounts


import (

	//  "context"
	//  "os"
	//  "time"
	//  "log"
	"github.com/StarGazer500/Department-Management-WebApp/internals/handlers/accounts"
	"github.com/StarGazer500/Department-Management-WebApp/internals/middlewares"
	// "fmt"

	
	"github.com/gin-gonic/gin"
)
func CreateUserAccountRoute(account_group *gin.RouterGroup){
	account_group.POST("/create-user", accounts.CreateUserAccountController)
}


func LoginRoute(account_group *gin.RouterGroup){
	account_group.POST("/login", accounts.LoginController)
	
}

func CreateRoleRoute(account_group *gin.RouterGroup){
	account_group.POST("/create-role", accounts.CreateRoleController)
	
}

func CheckIfuserIsValidRoute(account_group *gin.RouterGroup){
	account_group.GET("/is-user-valid",middlewares.AuthMiddleware(), accounts.CheckIfuserIsValidController)
	
}



