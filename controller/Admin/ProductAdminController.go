package Admin

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"github.com/gin-gonic/gin"
	"time"
)

func ProductAdminCreate(c *gin.Context) {
	var productCreate model.ProductAdminCreate
	var product model.Product
	var category model.Category
	c.ShouldBindJSON(&productCreate)
	initializers.DB.First(&model.Category{}, productCreate.CategoryId).Scan(&category)
	if category.ID == 0 {
		config.CustomResponse(c, 200, "not found category id", nil)
		return
	}
	product.ProductName = productCreate.ProductName
	product.CreateDate = time.Now()
	product.Status = true
	product.CategoryId = category.ID
	product.Image = productCreate.Image
	product.Description = productCreate.Description
	initializers.DB.Save(&product)
	config.CustomResponse(c, 200, "success", product)
	return
}
