package database

import (
	"fmt"
	"goTestProj/config"
	"goTestProj/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dsn = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
	config.Env.Get("database.host"),
	config.Env.Get("database.user"),
	config.Env.Get("database.password"),
	config.Env.Get("database.name"),
	config.Env.GetInt("database.port"),
)

var postgreCon = postgres.New(
	postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	},
)

func connectDB() *gorm.DB {
	db, err := gorm.Open(postgreCon, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "public.",
		},
	})

	if err != nil {
		fmt.Println("Connect Db failed : ", err)
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Region{})
	db.AutoMigrate(&models.Area{})

	return db.Session(&gorm.Session{
		PrepareStmt: true,
		// TableName: config.Env.Get("database.schemaName") + ".",
	})
}

var DB *gorm.DB

func DBinit() {
	DB = connectDB()
}
