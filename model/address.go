package model

import "time"

type Address struct {
	ID            int
	AddressDetail string
	ProvinceId    int
	DistrictId    int
	WardId        int
	AccountId     int
	IsDeleted     bool
	IsActive      bool
	CreateDate    time.Time
	IsDefault     bool
	PhoneNumber   string
	FullName      string
}