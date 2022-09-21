package model

import "time"

type Product struct {
	ID          int
	ProductName string
	Description string
	Image       string
	VoteAverage float64
	CreateDate  time.Time

	CategoryId int
	Status     bool
}
