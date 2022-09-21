package model

import "time"

type MaterialProduct struct {
	MaterialId int
	ProductId  int
	Status     bool
	CreateDate time.Time
}
