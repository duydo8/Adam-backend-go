package model

import "time"

type Product struct {
	ID          int
	ProductName string
	Description string
	Image       string
	VoteAverage float64
	CreateDate  time.Time `gorm:"column:create_date"`

	CategoryId      int
	Status          bool
	TagProducts     []TagProduct      `json:"-" gorm:"foreignKey:product_id"`
	Comment         []Comments        `json:"-" gorm:"foreignKey:product_id"`
	Favorite        []Favorite        `json:"-" gorm:"foreignKey:product_id"`
	DetailProduct   []DetailProduct   `json:"-" gorm:"foreignKey:product_id"`
	MaterialProduct []MaterialProduct `json:"-" gorm:"foreignKey:product_id"`
}
