package model

type Ward struct {
	ID         int
	Name       string
	Prefix     string
	ProvinceId int       `gorm:"column:province_id"`
	DistrictId int       `gorm:"column:district_id"`
	Addresses  []Address `json:"-" gorm:"foreignKey:address_id"`
}
