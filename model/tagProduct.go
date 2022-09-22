package model

import "time"

type TagProduct struct {
	TagId      int `gorm:"column:tag_id,primaryKey"`
	ProductId  int `gorm:"column:product_id,primaryKey"`
	Status     bool
	CreateDate time.Time `gorm:"column:create_date"`
}
