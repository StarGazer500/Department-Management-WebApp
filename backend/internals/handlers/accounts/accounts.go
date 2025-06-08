package accounts

import (
	"github.com/StarGazer500/Department-Management-WebApp/internals/services/accounts"

	// "database/sql"
	// "fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/lib/pq"
	// "golang.org/x/crypto/bcrypt"
)

func CreateUserAccountController(ctx *gin.Context) {
	var userAccount accounts.UserAccount

	if ctx.Request.Method == http.MethodPost {

		if err := ctx.ShouldBindJSON(&userAccount); err != nil {
			// If there's a binding error, return a 400 response and stop executionf
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Something Unexpected Happened",
			})
			return
		}

		message := accounts.CreateUserService(&userAccount)
		if message == "User not created" {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "account not created!",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Account Created Successfully!",
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": "account creation failed!",
	})

}

func LoginController(ctx *gin.Context) {
	var login accounts.LoginStruct

	if ctx.Request.Method == http.MethodPost {

		if err := ctx.ShouldBindJSON(&login); err != nil {
			// If there's a binding error, return a 400 response and stop executionf
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Something Unexpected Happened",
			})
			return
		}

		message := accounts.LoginService(&login, ctx)

		if message == "Username not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "User Not Found",
			})
			return

		}

		if message == "Invalid User" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid User",
			})
			return

		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": message,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "login failed!",
	})

}

func CreateRoleController(ctx *gin.Context) {
	var role accounts.RoleStruct

	if ctx.Request.Method == http.MethodPost {

		if err := ctx.ShouldBindJSON(&role); err != nil {
			// If there's a binding error, return a 400 response and stop executionf
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Something Unexpected Happened",
			})
			return
		}

		message := accounts.CreateRoleService(&role)
		if message == "User not created" {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "account not created!",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Role Created Successfully!",
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": "role creation failed!",
	})

}

func CheckIfuserIsValidController(ctx *gin.Context) {

	if ctx.Request.Method == http.MethodGet {

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User is Logi Session Is Still Valid!",
		})
		return
	}

}
