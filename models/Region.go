package models

import (
	"gorm.io/gorm"
)

type Region struct {
	gorm.Model
	RegionId int `gorm:"column:region_id;comment:id;uniqueKey"`
	// RegionTestColumn string `gorm:"column:region_test_column"`
	RegionName  string `gorm:"column:region_name;comment:處名稱"`
	RegionCode  int    `gorm:"column:region_code;comment:處代碼"`
	PhoneNumber string `gorm:"column:phone_number;comment:電話"`
	Status      bool   `gorm:"column:status;comment:狀態"`
	AreaCode    int    `gorm:"column:area_code;comment:母項目"`
}

// NewRegion 是一個建構函數，用於初始化 Region 結構體
func NewRegion(regionId int, regionName string, regionCode int, phoneNumber string, status bool, areaCode int) *Region {
	return &Region{
		RegionId:    regionId,
		RegionName:  regionName,
		RegionCode:  regionCode,
		PhoneNumber: phoneNumber,
		Status:      status,
		AreaCode:    areaCode,
	}
}
