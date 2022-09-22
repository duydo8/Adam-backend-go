package model

type Province struct {
	ID        int
	Name      string
	Prefix    string
	Wards     []Ward     `json:"-" gorm:"foreignKey:province_id"`
	Districts []District `json:"-" gorm:"foreignKey:province_id"`
	Addresses []Address  `json:"-" gorm:"foreignKey:province_id"`
}
