package model

import "time"

type Material struct {
	ID           int
	MaterialName string

	CreateDate      time.Time
	DetailProductId int
	Status          bool
}
