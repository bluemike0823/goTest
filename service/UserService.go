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
	db := database.DB
	users := []models.User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	db := database.DB
	user := models.User{}
	// db.Table("goTest.user").AutoMigrate(&user)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "ERROR : "+err.Error())
	}
	// fmt.Print(&user)
	// db.AutoMigrate(&user)
	err2 := db.Create(&user)
	if err2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {

	db := database.DB
	users := []models.User{}
	db.Find(&users)
	userId := c.Param("id")

	result := db.Where("user_id = ?", userId).Delete(users)

	if err := result.Error; err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	} else if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		// 表示沒有匹配的記錄被刪除
		fmt.Println("沒有符合條件的資料被刪除")
	} else {
		// 刪除成功
		fmt.Printf("成功刪除 %d 條資料\n", rowsAffected)
	}
	c.JSON(http.StatusOK, "user deleted")
}

func PutUser(c *gin.Context) {

	db := database.DB
	users := []models.User{}
	db.Find(&users)
	userId := c.Param("id")
	fmt.Println("userId : ", userId)
	newUserData := models.User{}
	err := c.ShouldBindJSON(&newUserData)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}

	result := db.Model(&models.User{}).Where("user_id = ?", userId).UpdateColumns(newUserData)
	if err := result.Error; err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	// result := db.Where("user_id = ?", userId).Update(users)

	updatedUser := models.User{}
	db.Where("user_id = ?", userId).First(&updatedUser)

	c.JSON(http.StatusOK, updatedUser)
}
