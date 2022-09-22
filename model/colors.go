package model

import "time"

type Color struct {
	ID        int
	ColorName string `gorm:"column:color_name"`

	CreateDate time.Time `gorm:"column:create_date"`

	Status        bool
	DetailProduct []DetailProduct `json:"-",gorm:"foreignKey:color_id;"`
}
