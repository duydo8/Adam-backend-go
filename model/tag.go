package model

import "time"

type Tag struct {
	ID              int
	TagName         string
	CreateDate      time.Time
	DetailProductId int
	Status          bool
}
