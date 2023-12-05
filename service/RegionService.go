package service

import (
	"goTestProj/database"
	"goTestProj/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var mockRegionsList = []models.Region{
	*models.NewRegion(1, "Example Region 1", 123, "123456789", true, 456),
	*models.NewRegion(2, "Example Region 2", 456, "987654321", false, 789),
	*models.NewRegion(3, "Example Region 3", 789, "555555555", true, 123),
	*models.NewRegion(4, "Example Region 4", 987, "999999999", false, 321),
}

func FindAllRegion(c *gin.Context) {

	db := database.DB
	region := []models.Region{}
	db.Find(&region)
	c.JSON(http.StatusOK, region)
}

func SetRegion(c *gin.Context) {
	db := database.DB
	region := models.Region{}

	err := c.ShouldBindJSON(&region)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "ERROR : "+err.Error())
	}

	err2 := db.Create(&region)
	if err2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, region)
}

func FindAllArea(c *gin.Context) {

	db := database.DB
	areas := []models.Area{}
	db.Find(&areas)
	c.JSON(http.StatusOK, areas)
}

func SetArea(c *gin.Context) {

	db := database.DB
	area := models.Area{}

	err := c.ShouldBindJSON(&area)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "ERROR : "+err.Error())
	}

	err2 := db.Create(&area)
	if err2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, area)
}
