package service

import (
	"fmt"
	"goTestProj/database"
	"goTestProj/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userList = []models.User{}

func FindAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

func PostUser(c *gin.Context) {
	db := database.DB
	user := models.User{}
	db.Table("goTest.user").AutoMigrate(&user)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "ERROR : "+err.Error())
	}
	fmt.Print(&user)
	db.AutoMigrate(&user)
	err2 := db.Create(&user)
	if err2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	// userId, _ := strconv.Atoi(c.Param("id"))
	// for index, user := range userList {
	// 	if userId == user.Id {
	// 		userList = append(userList[:index], userList[index+1:]...)
	// 		c.JSON(http.StatusOK, "user deleted")
	// 		return
	// 	}
	// }
	c.JSON(http.StatusGone, "Error")
}

func PutUser(c *gin.Context) {
	// updateUser := models.User{}
	// err := c.BindJSON(&updateUser)
	// if err != nil {
	// 	c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
	// }
	// userId, _ := strconv.Atoi(c.Param("id"))
	// for index, user := range userList {
	// 	if userId == user.Id {
	// 		userList[index] = updateUser
	// 		c.JSON(http.StatusOK, "user updated")
	// 		return
	// 	}
	// }
	c.JSON(http.StatusNotFound, "Error")
}
