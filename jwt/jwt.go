package jwt

import (
	"fmt"
	"goTestProj/error"
	"goTestProj/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		fmt.Println("== JWT() gin.HandlerFunc ==")

		code = error.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = error.INVALID_PARAMS
		} else {
			// claims, err := service.ParseToken(token)
			_, err := service.ParseToken(token)
			if err != nil {
				code = error.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
			//  else if time.Now().Unix() > claims.ExpiresAt {
			// 	code = error.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			// }
		}

		if code != error.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  error.GetMsg(code),
				"data": data,
			})

			fmt.Println("== JWT() : failed ==")

			c.Abort()
			return
		}

		fmt.Println("== JWT() : success ==")

		c.Next()

	}
}
