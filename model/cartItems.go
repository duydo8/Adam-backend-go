package model

import "time"

type CartItems struct {
	ID              int
	Quantity        int
	TotalPrice      float64
	AccountID       int `gorm:"column:account_id"`
	DetailProductID int `gorm:"column:detail_product_id"`
	Status          bool
	CreateDate      time.Time
	OrderID         int `gorm:"column:order_id"`
}
