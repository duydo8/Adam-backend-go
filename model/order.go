package model

import "time"

type Order struct {
	ID          int
	Status      int
	CreateDate  time.Time `gorm:"column:create_date"`
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
	HistoryOrder     []HistoryOrder `json:"-",gorm:"foreignKey:o;reference:ID"`
	CartItems        []CartItems    `json:"-",gorm:"foreignKey:OrderID;reference:ID"`
	DetailOrder      []DetailOrder  `json:"-",gorm:"foreignKey:OrderID;reference:ID"`
}
