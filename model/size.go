package model

import "time"

type Size struct {
	ID              int
	SizeName        string
	CreateDate      time.Time `gorm:"column:create_date"`
	DetailProductId int
	Status          bool
	DetailProducts  []DetailProduct `json:"-" gorm:"foreignKey:size_id"`
}
type SizeCreate struct {
	ID       int
	SizeName string
}
