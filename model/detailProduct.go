package model

import "time"

type DetailProduct struct {
	ID           int
	Quantity     int
	PriceImport  float64
	PriceExport  float64
	ProductImage string
	Status       bool
	ProductID    int
	ColorID      int
	SizeID       int
	CreateDate   time.Time   `gorm:"column:create_date"`
	CartItems    []CartItems `json:"-" gorm:"foreignKey:detail_product_id"`
}
