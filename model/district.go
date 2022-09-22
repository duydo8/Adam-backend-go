package model

type District struct {
	ID         int
	Name       string
	Prefix     string
	ProvinceID int
	Addresses  []Address `json:"-" gorm:"foreignKey:AddressID"`
	Ward       []Ward    `json:"-" gorm:"foreignKey:WardID"`
}
