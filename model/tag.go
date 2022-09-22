package model

import "time"

type Tag struct {
	ID              int
	TagName         string
	CreateDate      time.Time `gorm:"column:create_date"`
	DetailProductId int
	Status          bool
	TagProducts     []TagProduct `json:"-",gorm:"foreignKey:tag_id"`
}
