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
	HistoryOrders    []HistoryOrder `json:"-" gorm:"foreignKey:order_id"`
	CartItems        []CartItems    `json:"-" gorm:"foreignKey:order_id"`
	DetailOrders     []DetailOrder  `json:"-" gorm:"foreignKey:order_id"`
}
