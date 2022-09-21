package model

type Province struct {
	ID         int
	Name       string
	Prefix     string
	DistrictId int
	WardId     int
	AddressId  int
}
