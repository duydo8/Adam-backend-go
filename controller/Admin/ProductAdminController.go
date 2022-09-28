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
func ProductAdminUpdate(c *gin.Context) {
	var tagProduct model.TagProduct
	var tagProducts []model.TagProduct
	var tagProductsUpdate []model.TagProduct
	var materialProducts []model.MaterialProduct
	var materialProductsUpdate []model.MaterialProduct
	var materialProduct model.MaterialProduct
	var product model.Product
	var category model.Category
	var tags []model.Tag
	var materials []model.Material
	var productUpdate model.ProductUpdateAdmin
	c.ShouldBindJSON(&productUpdate)

	initializers.DB.First(&model.Product{}, productUpdate.ID).Scan(&product)
	initializers.DB.First(&model.Category{}, productUpdate.CategoryId).Scan(&category)
	if product.ID == 0 {
		config.CustomResponse(c, 200, "not found product", nil)
		return
	}
	if category.ID == 0 {
		config.CustomResponse(c, 200, "not found category", nil)
		return
	}
	for i := 0; i < len(productUpdate.TagIds); i++ {
		initializers.DB.First(&model.Tag{}, productUpdate.TagIds).Scan(&tags)
	}

	for i := 0; i < len(tags); i++ {

		tagProduct.TagId = tags[i].ID
		tagProduct.ProductId = productUpdate.ID
		tagProduct.CreateDate = time.Now()
		tagProduct.Status = true
		tagProductsUpdate = append(tagProductsUpdate, tagProduct)
		initializers.DB.Save(&tagProduct)
		initializers.DB.Raw("select * from tag_products where product_id=?", productUpdate.ID).Scan(&tagProducts)
		initializers.DB.Delete(&model.TagProduct{}, tagProducts[i])

	}
	for i := 0; i < len(productUpdate.MaterialIds); i++ {
		initializers.DB.First(&model.Material{}, productUpdate.MaterialIds).Scan(&materials)
	}

	for i := 0; i < len(materials); i++ {

		materialProduct.MaterialId = materials[i].ID
		materialProduct.ProductId = productUpdate.ID
		materialProduct.CreateDate = time.Now()
		materialProduct.Status = true
		materialProductsUpdate = append(materialProductsUpdate, materialProduct)
		initializers.DB.Save(&materialProduct)
		initializers.DB.Raw("select * from material_products where product_id=?", productUpdate.ID).Scan(&materialProducts)
		initializers.DB.Delete(&model.MaterialProduct{}, materialProducts[i])

	}
	if len(materialProductsUpdate) != 0 {
		product.MaterialProducts = materialProductsUpdate
	}
	if len(tagProductsUpdate) != 0 {
		product.TagProducts = tagProductsUpdate
	}
	product.Status = productUpdate.Status
	if productUpdate.CategoryId != 0 {
		product.CategoryId = productUpdate.CategoryId
	}
	if productUpdate.Description != "" {
		product.Description = productUpdate.Description
	}
	if productUpdate.Image != "" {
		product.Image = productUpdate.Image
	}
	if productUpdate.ProductName != "" {
		product.ProductName = productUpdate.ProductName
	}
	initializers.DB.Save(&product)
	config.CustomResponse(c, 200, "success", product)

}
func ProductAdminDelete(c *gin.Context) {
	var product model.Product
	id := c.Query("id")
	initializers.DB.First(&model.Tag{}, id).Scan(&product)
	if product.ID != 0 {
		product.Status = false
		initializers.DB.Save(&product)
		config.CustomResponse(c, 200, "success", product)
		return
	} else {
		config.CustomResponse(c, 200, "not found id", nil)
		return
	}
}
func ProductCreateArrayOption(c *gin.Context) {
	var productAdminCreateArray model.ProductAdminCreateArray
	c.ShouldBindJSON(&productAdminCreateArray)
	var category model.Category
	var tags []model.Tag
	var materials []model.Material

	var product model.Product
	var tagProduct model.TagProduct
	var materialProduct model.MaterialProduct
	var materialProducts []model.MaterialProduct
	var tagProducts []model.TagProduct
	initializers.DB.First(&model.Category{}, productAdminCreateArray.CategoryId).Scan(&category)

	if category.ID == 0 {
		config.CustomResponse(c, 200, "not found category", nil)
		return
	}

	product.ProductName = productAdminCreateArray.ProductName
	product.CategoryId = productAdminCreateArray.CategoryId
	product.Description = productAdminCreateArray.Description
	product.Image = productAdminCreateArray.Image
	initializers.DB.Save(&product)
	for i := 0; i < len(productAdminCreateArray.TagIdList); i++ {
		initializers.DB.First(&model.Tag{}, productAdminCreateArray.TagIdList[i]).Scan(&tags)

	}
	for i := 0; i < len(tags); i++ {
		tagProduct.TagId = tags[i].ID
		tagProduct.ProductId = product.ID
		tagProduct.CreateDate = time.Now()
		tagProduct.Status = true
		initializers.DB.Save(&tagProduct)
		tagProducts = append(tagProducts, tagProduct)
	}
	for i := 0; i < len(productAdminCreateArray.MaterialIdList); i++ {
		initializers.DB.First(&model.Tag{}, productAdminCreateArray.MaterialIdList[i]).Scan(&materials)

	}
	for i := 0; i < len(materials); i++ {
		materialProduct.MaterialId = materials[i].ID
		materialProduct.ProductId = product.ID
		materialProduct.CreateDate = time.Now()
		materialProduct.Status = true
		materialProducts = append(materialProducts, materialProduct)

		initializers.DB.Save(&materialProduct)
	}
	product.MaterialProducts = materialProducts
	product.TagProducts = tagProducts
	initializers.DB.Save(&product)
	config.CustomResponse(c, 200, "success", nil)
	return

}
