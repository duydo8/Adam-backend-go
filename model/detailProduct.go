package model

import "time"

type detailProduct struct {
	ID           int
	Quantity     int
	PriceImport  float64
	PriceExport  float64
	ProductImage string
	Status       bool
	ProductId    int
	ColorId      int
	SizeId       int
	CreateDate   time.Time
}
