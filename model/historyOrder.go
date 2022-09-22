package model

import "time"

type HistoryOrder struct {
	ID int

	Description string
	UpdateTime  time.Time `gorm:"column:update_time"`
	TotalPrice  float64
	Status      int
	OrderId     int
}
