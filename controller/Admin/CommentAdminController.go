package Admin

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"github.com/gin-gonic/gin"
	"time"
)

func CommentAdminCreate(c *gin.Context) {
	var commentCreate model.CommentCreate
	var comment model.Comments
	var product model.Product
	c.ShouldBindJSON(&commentCreate)
	var account model.Accounts
	initializers.DB.First(&model.Accounts{}, commentCreate.AccountID).Scan(&account)
	if account.ID == 0 {
		config.CustomResponse(c, 200, "please login before comment", nil)
		return
	}
	initializers.DB.First(&model.Comments{}, commentCreate.CommentParentID).Scan(&comment)
	if comment.ID == 0 {
		config.CustomResponse(c, 200, "not found comment parent id", nil)
		return
	}
	initializers.DB.First(&model.Product{}, commentCreate.ProductID).Scan(&product)
	if product.ID == 0 {
		config.CustomResponse(c, 200, "not found product", nil)
		return
	}
	comment.CommentParentID = commentCreate.CommentParentID
	comment.Content = commentCreate.Content
	comment.Vote = commentCreate.Vote
	comment.AccountID = commentCreate.AccountID
	comment.CommentStatus = 1
	comment.ProductID = commentCreate.ProductID
	comment.CreateDate = time.Now()
	initializers.DB.Save(&comment)
	config.CustomResponse(c, 200, "success", comment)
	return
}
