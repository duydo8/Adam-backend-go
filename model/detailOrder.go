package model

import "time"

type DetailOrder struct {
	ID              int
	Quantity        int
	Price           float64
	IsDeleted       bool
	DetailOrderCode string
	Status          bool
	CreateDate      time.Time `gorm:"column:create_date"`
	Reason          string
	DetailProductID int
	OrderID         int
}
