package model

import "time"

type Product struct {
	ID          int
	ProductName string
	Description string
	Image       string
	VoteAverage float64
	CreateDate  time.Time `gorm:"column:create_date"`

	CategoryId       int
	Status           bool
	TagProducts      []TagProduct      `json:"-" gorm:"foreignKey:product_id"`
	Comments         []Comments        `json:"-" gorm:"foreignKey:product_id"`
	Favorites        []Favorite        `json:"-" gorm:"foreignKey:product_id"`
	DetailProducts   []DetailProduct   `json:"-" gorm:"foreignKey:product_id"`
	MaterialProducts []MaterialProduct `json:"-" gorm:"foreignKey:product_id"`
}
type ProductAdminCreate struct {
	ID          int
	ProductName string
	Description string
	Image       string
	CategoryId  int
}
type ProductUpdateAdmin struct {
	ID          int
	ProductName string
	Description string
	Image       string
	VoteAverage float64
	CategoryId  int
	TagIds      []int
	MaterialIds []int
	Status      bool
}
type ProductAdminCreateArray struct {
	ProductName    string
	CategoryId     int
	Description    string
	Image          string
	TagIdList      []int
	MaterialIdList []int
}
