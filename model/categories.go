package model

import "time"

type Category struct {
	ID           int
	CategoryName string

	CreateDate       time.Time
	CategoryParentId int
	Status           bool
}
