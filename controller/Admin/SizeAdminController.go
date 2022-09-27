package Admin

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"github.com/gin-gonic/gin"
	"time"
)

func SizeAdminFindAll(c *gin.Context) {
	var size []model.Size
	initializers.DB.Where("status =1").Find(&size)
	config.CustomResponse(c, 200, "success", size)
	return

}
func SizeAdminCreate(c *gin.Context) {
	var size model.SizeCreate
	var size1 model.Size
	c.ShouldBindJSON(&size)
	initializers.DB.Where("size_name=?", size.SizeName).First(&size1)
	if size1.ID != 0 {
		config.CustomResponse(c, 200, "duplicate Size name", nil)
		return
	}
	size1.CreateDate = time.Now()
	size1.SizeName = size.SizeName
	size1.Status = true
	initializers.DB.Save(&size1)
	config.CustomResponse(c, 200, "success", size1)
	return

}

func SizeAdminUpdate(c *gin.Context) {
	var size model.Size
	var size2 model.Size
	var size1 model.SizeCreate
	c.ShouldBindJSON(&size1)
	initializers.DB.Where("id=?", size1.ID).First(&size)
	if size1.ID == 0 {
		config.CustomResponse(c, 200, "not found", nil)
		return
	}
	initializers.DB.Where("size_name=?", size1.SizeName).First(&size2)
	if size2.ID != 0 {
		config.CustomResponse(c, 200, "duplicate size name", nil)
		return
	}
	size.SizeName = size1.SizeName

	initializers.DB.Save(&size)
	config.CustomResponse(c, 200, "success", size)
	return
}
func SizeAdminDelete(c *gin.Context) {
	var size model.Size
	id := c.Query("id")
	initializers.DB.First(&model.Size{}, id).Scan(&size)
	if size.ID != 0 {
		size.Status = false
		initializers.DB.Save(&size)
		config.CustomResponse(c, 200, "success", size)
		return
	} else {
		config.CustomResponse(c, 200, "not found id", nil)
		return
	}
}
