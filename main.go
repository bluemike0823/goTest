package main

import (
	"goTestProj/config"
	"goTestProj/service"
	"log"
	"net/http"
	"os"
	"time"

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

	log.Println(" : log created")

	router := gin.Default()
	http.HandleFunc("/", service.MethodHandler)
	api := router.Group("/api")
	route.AddAuthRouter(api)
	route.AddUserRouter(api)
	route.AddRegionRouter(api)

	database.DBinit()
	log.Println(" : service start")
	// router.Run(":8080")
	router.Run(":" + strconv.Itoa(config.Env.GetInt("server.port")))
}

func logInit() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)

	currentDate := time.Now().Format("2006-01-01")

	if _, err := os.Stat("log"); os.IsNotExist(err) {
		if err := os.Mkdir("log", 0755); err != nil {
			log.Fatal("cannot create log dir: ", err)
		}
	}

	file, err := os.OpenFile("log/logfile"+currentDate+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("cannot open log file : ", err)
	}

	log.SetOutput(file)

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)

}
