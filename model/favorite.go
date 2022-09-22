package model

import "time"

type Favorite struct {
	ProductId  int
	AccountId  int
	CreateDate time.Time `gorm:"column:create_date"`
	Status     bool
}
