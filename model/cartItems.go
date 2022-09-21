package model

import "time"

type CartItems struct {
	ID              int
	Quantity        int
	TotalPrice      float64
	AccountId       int
	DetailProductId int
	Status          bool
	CreateDate      time.Time
	OrderId         int
}
