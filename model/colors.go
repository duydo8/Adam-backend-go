package model

import "time"

type Color struct {
	ID        int
	ColorName string `gorm:"column:color_name"`

	CreateDate time.Time `gorm:"column:create_date"`

	Status         bool
	DetailProducts []DetailProduct `json:"-",gorm:"foreignKey:color_id;"`
}
type ColorCreate struct {
	ID        int
	ColorName string
}
