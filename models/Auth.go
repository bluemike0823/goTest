package models

// import "goTestProj/database"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"unique" json:"password"`
}

// func CheckAuth(username string, password string) bool {
// 	var auth Auth
// 	database.DB.Select("id").Where(Auth{
// 		Username: username,
// 		Password: password,
// 	}).First(&auth)

// 	if auth.ID > 0 {
// 		return true
// 	}

// 	return false

// }