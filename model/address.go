package model

import "time"

type Address struct {
	ID            int `gorm:"primaryKey,AUTO_INCREMENT"`
	AddressDetail string
	ProvinceID    int `gorm:"column:province_id"`
	DistrictID    int `gorm:"column:district_id"`
	WardID        int `gorm:"column:ward_id"`
	AccountID     int `gorm:"column:account_id"`
	Status        bool
	CreateDate    time.Time
	IsDefault     bool
	PhoneNumber   string
	FullName      string
}
