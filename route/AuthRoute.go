package route

import (
	"github.com/gin-gonic/gin"

	"goTestProj/service"
)

func AddAuthRouter(r *gin.RouterGroup) {
	region := r.Group("/auth")
	region.GET("/", service.GetAuth)

}

// func GetAuth(c *gin.Context) {
	
// 	username := c.Query("username")
// 	password := c.Query("password")

// 	valid := validation.Validation{}
// 	auth := auth{
// 		Username: username,
// 		Password: password,
// 	}
// 	result, _ := valid.Valid(&auth)

// 	data := make(map[string]interface{})
// 	code := error.INVALID_PARAMS

// 	if result {
// 		isExist := service.CheckAuth(username, password)
// 		if isExist {
// 			token, err := service.GenerateToken(username, password)
// 			if err != nil {
// 				code = error.ERROR_AUTH_TOKEN
// 			} else {
// 				data["token"] = token
// 				code = error.SUCCESS
// 			}
// 		} else {
// 			code = error.ERROR_AUTH
// 		}
// 	} else {
// 		for _, err := range valid.Errors {
// 			log.Println(err.Key, err.Message)
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code" : code,
// 		"msg" : error.GetMsg(code),
// 		"data" : data,
// 	})

// }