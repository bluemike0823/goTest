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

	fmt.Println("=== FindAllUser == : ")
	db := database.DB
	users := []models.User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	db := database.DB
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "ERROR : "+err.Error())
	}
	err2 := db.Create(&user)
	if err2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// func autoMigrateModels(db *gorm.DB) {
// 	models := []ModelInterface{
// 		&models.User{},
// 		&models.Region{},
// 	}
// }

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

	updatedUser := models.User{}
	db.Where("user_id = ?", userId).First(&updatedUser)

	c.JSON(http.StatusOK, updatedUser)
}

func FindAllJurisdiction(c *gin.Context) {

	fmt.Println("=== FindAllJurisdiction ==")
	userId := c.Param("id")
	var emptyList []string
	// userIdList := findJurisdictionByUserId(userId, emptyList)
	userIdList := findJurisdictionByUserId(userId, emptyList)

	c.JSON(http.StatusOK, userIdList)
}

func findJurisdictionByUserId(userId string, nowList []string) []string {

	fmt.Println("=== FindByUserId == : ", userId)
	db := database.DB
	result := []models.User{}
	db.Model(&models.User{}).Where("superior = ?", userId).Find(&result)

	if len(result) < 1 {

		fmt.Println("=== --> append == : ", userId)
		return append(nowList, userId)
	}

	for _, v := range result {
		nowList = append(nowList, userId)
		ansList := findJurisdictionByUserId(v.UserID, nowList)
		append(nowList, ansList)
	}

	return nowList
}
