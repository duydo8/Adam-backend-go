package model

type Ward struct {
	ID         int
	Name       string
	Prefix     string
	ProvinceID int       `gorm:"column:province_id"`
	DistrictID int       `gorm:"column:district_id"`
	Addresses  []Address `json:"-" gorm:"foreignKey:ward_id"`
}
