package database

import (
	"fmt"
	"goTestProj/config"
	"goTestProj/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// var dsn = fmt.Sprintf(
// 	"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
// 	config.Env.Get("database.host"),
// 	config.Env.Get("database.user"),
// 	config.Env.Get("database.password"),
// 	config.Env.Get("database.name"),
// 	config.Env.GetInt("database.port"),
// )

var dsn_mysql = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
	config.Env.Get("database.user"),
	config.Env.Get("database.password"),
	config.Env.Get("database.host"),
	config.Env.GetInt("database.port"),
	config.Env.Get("database.name"))

var postgreCon = postgres.New(
	postgres.Config{
		DSN:                  dsn_mysql,
		PreferSimpleProtocol: true,
	},
)

func connectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn_mysql), &gorm.Config{
		//  db, err := gorm.Open(postgreCon, &gorm.Config{
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
	db.AutoMigrate(&models.Auth{})
	// err = db.Create(&models.Auth{Username: "admin", Password: "admin"}).Error
	db.Create(&models.Auth{Username: "admin", Password: "admin"})
	// if err != nil {
	//     // panic("無法創建初始資料")
	// }

	return db.Session(&gorm.Session{
		PrepareStmt: true,
		// TableName: config.Env.Get("database.schemaName") + ".",
	})
}

var DB *gorm.DB

func DBinit() {
	DB = connectDB()
}
