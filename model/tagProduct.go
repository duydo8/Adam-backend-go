package model

import "time"

type TagProduct struct {
	TagId      int `gorm:"column:tag_id;primaryKey;autoIncrement:false"`
	ProductId  int `gorm:"column:product_id;primaryKey;autoIncrement:false"`
	Status     bool
	CreateDate time.Time `gorm:"column:create_date"`
}
