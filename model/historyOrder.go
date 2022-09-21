package model

import "time"

type HistoryOrder struct {
	ID int

	Description string
	UpdateTime  time.Time
	TotalPrice  float64
	Status      int
	OrderId     int
}

