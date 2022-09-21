package model

import "time"

type DetailOrder struct {
	ID              int
	Quantity        int
	Price           float64
	IsDeleted       bool
	DetailOrderCode string
	Status          bool
	CreateDate      time.Time
	Reason          string
	DetailProductId int
	OrderId         int
}
