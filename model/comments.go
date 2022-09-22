package model

import "time"

type Comments struct {
	ID              int
	Content         string
	Vote            int
	CreateDate      time.Time `gorm:"column:create_date"`
	CommentStatus   int
	CommentParentID int
	AccountID       int `gorm:"column:account_id"`
	ProductID       int `gorm:"column:product_id"`
	Status          bool
}
