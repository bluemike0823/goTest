package models

import (
	"gorm.io/gorm"
)

type Area struct {
	gorm.Model
	AreaId int `gorm:"column:area_id;unique"`
	// RegionTestColumn string `gorm:"column:region_test_column"`
	AreaName string `gorm:"column:area_name;unique"`
}

// NewRegion 是一個建構函數，用於初始化 Region 結構體
func NewArea(areaId int, areaName string) *Area {
	return &Area{
		AreaId:   areaId,
		AreaName: areaName,
	}
}
