package main

import (
	"goTestProj/config"

	"goTestProj/database"
	"goTestProj/route"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	api := router.Group("/api")
	route.AddUserRouter(api)

	database.DBinit()

	// router.Run(":8080")
	router.Run(":" + strconv.Itoa(config.Env.GetInt("server.port")))
}

// func test(c *gin.Context) {
// 	str := []byte("OK")
// 	c.Data(http.StatusOK, "test/plain", str)
// }

// func plus(c *gin.Context) {
// 	aStr := c.Param("a")
// 	bStr := c.Param("b")

// 	a, err := strconv.Atoi(aStr)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "a error",
// 		})
// 		return
// 	}

// 	b, err := strconv.Atoi(bStr)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "b error",
// 		})
// 		return
// 	}

// 	result := a + b

// 	c.JSON(200, gin.H{
// 		"ans": result,
// 	})
// }
