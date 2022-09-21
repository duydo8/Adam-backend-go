package model

import "time"

type Order struct {
	ID          int
	Status      int
	CreateDate  time.Time
	AccountId   int
	FullName    string
	PhoneNumber string
	AmountPrice float64

	SalePrice        float64
	TotalPrice       float64
	AddressId        int
	AddressDetail    string
	OrderCode        string
	ReturnOrderPrice float64
}
