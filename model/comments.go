package model

import "time"

type Comments struct {
	ID              int
	Content         string
	Vote            int
	CreateDate      time.Time
	CommentStatus   int
	CommentParentId int
	AccountId       int
	ProductId       int
	Status          bool
}
