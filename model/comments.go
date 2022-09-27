package model

import "time"

type Comments struct {
	ID              int
	Content         string
	Vote            int
	CreateDate      time.Time `gorm:"column:create_date"`
	CommentStatus   int
	CommentParentID int `gorm:"column:comment_parent_id"`
	AccountID       int `gorm:"column:account_id"`
	ProductID       int `gorm:"column:product_id"`
}
type CommentCreate struct {
	ID              int
	Content         string
	Vote            int
	CommentParentID int
	AccountID       int
	ProductID       int
}
