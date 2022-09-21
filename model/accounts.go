package model

import "time"

type Accounts struct {
	ID               int
	PhoneNumber      string `gorm:"column:phone_number,NOT NULL"`
	FullName         string `gorm:"column:full_name,NOT NULL"`
	Username         string `gorm:"column:user_name"`
	Password         string `gorm:"column:password"`
	Email            string `gorm:"column:not null"`
	Photo            string
	Role             bool
	VerificationCode string `gorm:"column:verification_code"`
	TimeValid        time.Time
	Status           bool
	CreateDate       time.Time
	Priority         float64
}
