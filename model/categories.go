package model

import "time"

type Category struct {
	ID           int
	CategoryName string `gorm:"column:category_name"`

	CreateDate       time.Time `gorm:"column:create_date"`
	CategoryParentID int       `gorm:"column:category_parent_id"`
	Status           bool
	Products         []Product `json:"-",gorm:"foreignKey:category_id"`
}
type CategoryAdminCreate struct {
	CategoryName     string
	CategoryParentID int
}
