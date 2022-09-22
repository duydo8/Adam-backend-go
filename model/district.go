package model

type District struct {
	ID         int
	Name       string
	Prefix     string
	ProvinceID int       `gorm:"column:province_id"`
	Addresses  []Address `json:"-" gorm:"foreignKey:district_id"`
	Wards      []Ward    `json:"-" gorm:"foreignKey:district_id"`
}
