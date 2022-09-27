package Admin

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"github.com/gin-gonic/gin"
	"time"
)

func AdminCategoryCreate(context *gin.Context) {
	var categoryCreate model.CategoryAdminCreate
	var category model.Category
	var cate model.Category
	context.ShouldBindJSON(&categoryCreate)
	if categoryCreate.CategoryParentID == 0 {

	} else if initializers.DB.First(&model.Category{}, categoryCreate.CategoryParentID).Scan(&cate); cate.ID == 0 {

		config.CustomResponse(context, 200, "not found category", nil)
		return
	}
	category.CategoryName = categoryCreate.CategoryName
	category.CreateDate = time.Now()
	category.Status = true
	category.CategoryParentID = categoryCreate.CategoryParentID
	initializers.DB.Save(&category)
	config.CustomResponse(context, 200, "success", category)
	return
}
func AdminCategoryFindAll(c *gin.Context) {
	var category []model.Category
	initializers.DB.Find(&category)
	config.CustomResponse(c, 200, "success", category)
}
func AdminCategoryFindAllCategoryParent(c *gin.Context) {
	var category []model.Category
	initializers.DB.Raw("select * from categories where category_parent_id=?", 0).Scan(&category)
	config.CustomResponse(c, 200, "success", category)
}
func AdminCategoryUpdate(c *gin.Context) {
	var category model.Category
	id := c.Query("id")
	cateName := c.Query("category_name")
	initializers.DB.Where("id = ?", id).Find(&model.Category{}).Scan(&category)
	if category.ID == 0 {
		config.CustomResponse(c, 200, "not found", nil)
		return

	}
	category.CategoryName = cateName
	initializers.DB.Save(&category)
	config.CustomResponse(c, 200, "success", category)
	return
}
func AdminCategoryDeleteByID(c *gin.Context) {
	id := c.Query("id")
	err := initializers.DB.Delete(&model.Category{}, id)
	if err != nil {
		config.CustomResponse(c, 200, "not found", nil)
		return
	}
	config.CustomResponse(c, 200, "success", nil)
	return
}
func AdminCategoryFindCategoryByParentId(c *gin.Context) {
	id := c.Query("id")
	var category []model.Category
	initializers.DB.Raw("select * from categories where category_parent_id=?", id).Scan(&category)

	config.CustomResponse(c, 200, "success", category)
	return
}
