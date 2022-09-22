package model

import "time"

type DetailProduct struct {
	ID           int
	Quantity     int
	PriceImport  float64
	PriceExport  float64
	ProductImage string
	Status       bool
	ProductID    int         `gorm:"column:product_id"`
	ColorID      int         `gorm:"column:color_id"`
	SizeID       int         `gorm:"column:size_id"`
	CreateDate   time.Time   `gorm:"column:create_date"`
	CartItems    []CartItems `json:"-" gorm:"foreignKey:detail_product_id"`
}
