package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      string    `gorm:"column:user_id;uniqueKey"`
	UserName    string    `gorm:"column:user_name"`
	RegionCode  int       `gorm:"column:region_code"`
	Gender      string    `gorm:"column:gender"`
	PhoneNumber string    `gorm:"column:phone_number"`
	JobTitle    string    `gorm:"column:job_title"`
	JoinDate    time.Time `gorm:"column:join_date"`
	Status      bool      `gorm:"column:status"`
	Superior    string    `gorm:"column:superior"`
}
