package model

import "time"

type DiscountOrder struct {
	ID            int
	DiscountName  string
	Description   string
	CreateDate    time.Time
	StartDate     time.Time
	EndDate       time.Time
	Status        bool
	SalePrice     float64
	OrderMinPrice float64
	OrderMaxPrice float64
	EventId       int
}
