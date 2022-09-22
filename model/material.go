package model

import "time"

type Material struct {
	ID           int
	MaterialName string

	CreateDate       time.Time `gorm:"column:create_date"`
	DetailProductId  int
	Status           bool
	MaterialProducts []MaterialProduct `json:"-" gorm:"foreignKey:material_id"`
}
