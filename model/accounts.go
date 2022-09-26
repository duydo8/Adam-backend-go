package model

import "time"

type Accounts struct {
	ID               int    `gorm:"primaryKey,AUTO_INCREMENT"`
	PhoneNumber      string ` gorm:"column:phone_number"`
	FullName         string `json:"FullName" gorm:"column:full_name"`
	Username         string `gorm:"column:username"`
	Password         string
	Email            string
	Photo            string
	Role             bool
	VerificationCode string `gorm:"column:verification_code"`
	TimeValid        time.Time
	Status           bool
	CreateDate       time.Time
	Priority         float64
	Addresses        []Address   `json:"-" gorm:"foreignKey:account_id"`
	Comments         []Comments  `json:"-" gorm:"foreignKey:account_id"`
	Favorites        []Favorite  `json:"-" gorm:"foreignKey:account_id"`
	CartItems        []CartItems `json:"-" gorm:"foreignKey:account_id"`
	Orders           []Order     `json:"-" gorm:"foreignKey:account_id"`
}
