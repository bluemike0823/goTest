package service

import (
	"fmt"
	"goTestProj/database"
	"goTestProj/error"
	"goTestProj/models"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println("=== username : ", username)
	fmt.Println("=== password : ", password)

	valid := validation.Validation{}
	auth := auth{
		Username: username,
		Password: password,
	}
	result, _ := valid.Valid(&auth)

	data := make(map[string]interface{})
	code := error.INVALID_PARAMS

	if result {
		isExist := CheckAuth(username, password)
		if isExist {
			token, err := GenerateToken(username, password)
			if err != nil {
				code = error.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = error.SUCCESS
			}
		} else {
			code = error.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : error.GetMsg(code),
		"data" : data,
	})

}

func CheckAuth(username string, password string) bool {
	var auth models.Auth
	database.DB.Select("id").Where(models.Auth{
		Username: username,
		Password: password,
	}).First(&auth)

	if auth.ID > 0 {
		return true
	}

	return false

}