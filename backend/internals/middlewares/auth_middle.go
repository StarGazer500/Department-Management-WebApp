package middlewares

import (
	"fmt"

	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/StarGazer500/Department-Management-WebApp/internals/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Check if user is authenticated using a cookie
		access_token, err := c.Cookie("access")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"success":false,"message": "no access token or token expired"})
			c.Abort() // Stop further processing of the request
			return
		}

		verifyaccess, err := utils.VerifyAccessToken(access_token, []byte(os.Getenv("JWT_SECRET_KEY")))
		fmt.Println("access verified",verifyaccess)
		if err != nil {

			fmt.Println("err", verifyaccess)
			// If user cookie is not set, return a 401 Unauthorized error
			refresh_token, err := c.Cookie("refresh")
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"success":false,"message": "user access session expired"})
				c.Abort() // Stop further processing of the request
				return
			}

			verifyrefresh, err := utils.VerifyRefreshToken(refresh_token, []byte(os.Getenv("JWT_SECRET_KEY")))
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"success":false,"message": "user access session expired"})
				c.Abort() // Stop further processing of the request
				return

			}
			claim:=utils.TokenClaimStruct{
				MyAuthServer:    verifyrefresh["iss"].(string),
				AuthUserEmail:   verifyrefresh["sub"].(string),
				AuthUserId:      0,
	         } 

			 if id, ok := verifyrefresh["id"].(float64); ok {
    			claim.AuthUserId = int(id)
			 }
			newaccess, err := utils.GenerateAccessToken(claim)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"success":false,"message": "no claim payload"})
				c.Abort() // Stop further processing of the request
			}
			fmt.Println("new access",newaccess)
			c.SetCookie("access", newaccess, 3600, "/", "127.0.0.1", false, true)
			fmt.Println("newaccess", newaccess)
			fmt.Println("newaccess", verifyrefresh)
			c.Next()

		}

		c.Next()

		fmt.Println(verifyaccess)
	}

	// If the user is authenticated, proceed to the next handler

}