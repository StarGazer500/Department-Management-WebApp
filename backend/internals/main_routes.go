package internals



import (
	// "github.com/StarGazer500/ayigya/controllers"
	// "github.com/StarGazer500/ayigya/middlewares"
 	accounts "github.com/StarGazer500/Department-Management-WebApp/internals/accounts/routes"

	"github.com/gin-gonic/gin"
)



func RegisterAllRoutes(engine *gin.Engine){

	account_group := engine.Group("/account")
	accounts.RegisterAccountRoutes(account_group)
	// accounts.LoginRoute(account_group)
	// accounts.CreateRoleRoute(account_group)
	// accounts.CreateUserAccountRoute(account_group)
	// accounts.CheckIfuserIsValidRoute(account_group)

}