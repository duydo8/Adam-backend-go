package model

import "time"

type Size struct {
	ID              int
	SizeName        string
	CreateDate      time.Time
	DetailProductId int
	Status          bool
}
