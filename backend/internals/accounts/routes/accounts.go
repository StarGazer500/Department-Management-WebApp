package routes


import (

	//  "context"
	//  "os"
	//  "time"
	//  "log"
	"github.com/StarGazer500/Department-Management-WebApp/internals/accounts/handlers"
	"github.com/StarGazer500/Department-Management-WebApp/internals/middlewares"
	// "fmt"

	
	"github.com/gin-gonic/gin"
)
func CreateUserAccountRoute(account_group *gin.RouterGroup){
	account_group.POST("/create-user", handlers.CreateUserAccountController)
}


func LoginRoute(account_group *gin.RouterGroup){
	account_group.POST("/login", handlers.LoginController)
	
}

func CreateRoleRoute(account_group *gin.RouterGroup){
	account_group.POST("/create-role", handlers.CreateRoleController)
	
}

func CheckIfuserIsValidRoute(account_group *gin.RouterGroup){
	account_group.GET("/is-user-valid",middlewares.AuthMiddleware(), handlers.CheckIfuserIsValidController)
	
}




func RegisterAccountRoutes(account_group *gin.RouterGroup){
	// account_group := engine.Group("/account")
	LoginRoute(account_group)
	CreateRoleRoute(account_group)
	CreateUserAccountRoute(account_group)
	CheckIfuserIsValidRoute(account_group)

}



