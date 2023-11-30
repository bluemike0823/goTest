package main

import (
	"goTestProj/config"
	"log"
	"os"

	"goTestProj/database"
	"goTestProj/route"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func main() {

	logInit()

	router := gin.Default()
	api := router.Group("/api")
	route.AddUserRouter(api)

	database.DBinit()

	// router.Run(":8080")
	router.Run(":" + strconv.Itoa(config.Env.GetInt("server.port")))
}

func logInit() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)

	if err := os.Mkdir("log", 0755); err != nil {
		log.Fatal("cannot create log dir: ", err)
	}

	file, err := os.OpenFile("log/logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("cannot open log file : ", err)
	}

	log.SetOutput(file)

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)

}
