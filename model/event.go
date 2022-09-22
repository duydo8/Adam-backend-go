package model

import "time"

type Event struct {
	ID             int
	EventName      string
	Description    string
	CreateDate     time.Time `gorm:"column:create_date"`
	StartDate      time.Time
	EndDate        time.Time
	Status         bool
	Type           bool
	Image          string
	DiscountOrders []DiscountOrder `json:"-" gorm:"foreignKey:event_id"`
}
