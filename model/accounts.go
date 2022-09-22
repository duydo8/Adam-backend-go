package model

import "time"

type Accounts struct {
	ID               int    `gorm:"primaryKey,AUTO_INCREMENT"`
	PhoneNumber      string `gorm:"column:phone_number,NOT NULL"`
	FullName         string `gorm:"column:full_name,NOT NULL"`
	Username         string `gorm:"column:user_name"`
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
