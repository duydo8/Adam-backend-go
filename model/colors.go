package model

import "time"

type Color struct {
	ID        int
	ColorName string

	CreateDate      time.Time
	DetailProductId int
	Status          bool
}
