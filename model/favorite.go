package model

import "time"

type Favorite struct {
	ProductId  int
	AccountId  int
	CreateDate time.Time
	Status     bool
}
