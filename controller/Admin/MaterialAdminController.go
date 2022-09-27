package Admin

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"github.com/gin-gonic/gin"
	"time"
)

func MaterialAdminFindAll(c *gin.Context) {
	var tag []model.Material
	initializers.DB.Where("status =1").Find(&tag)
	config.CustomResponse(c, 200, "success", tag)
	return

}
func MaterialAdminCreate(c *gin.Context) {
	var material model.MaterialCreate
	var material1 model.Material
	c.ShouldBindJSON(&material)
	initializers.DB.Where("material_name=?", material.MaterialName).First(&material1)
	if material1.ID != 0 {
		config.CustomResponse(c, 200, "duplicate Material name", nil)
		return
	}
	material1.CreateDate = time.Now()
	material1.MaterialName = material.MaterialName
	material1.Status = true
	initializers.DB.Save(&material1)
	config.CustomResponse(c, 200, "success", material1)
	return

}

func MaterialAdminUpdate(c *gin.Context) {
	var material model.Material
	var material2 model.Material
	var material1 model.MaterialCreate
	c.ShouldBindJSON(&material1)
	initializers.DB.Where("id=?", material1.ID).First(&material)
	if material1.ID == 0 {
		config.CustomResponse(c, 200, "not found", nil)
		return
	}
	initializers.DB.Where("material_name=?", material1.MaterialName).First(&material2)
	if material2.ID != 0 {
		config.CustomResponse(c, 200, "duplicate Material name", nil)
		return
	}
	material.MaterialName = material1.MaterialName

	initializers.DB.Save(&material)
	config.CustomResponse(c, 200, "success", material)
	return
}
func MaterialAdminDelete(c *gin.Context) {
	var material model.Material
	id := c.Query("id")
	initializers.DB.First(&model.Material{}, id).Scan(&material)
	if material.ID != 0 {
		material.Status = false
		initializers.DB.Save(&material)
		config.CustomResponse(c, 200, "success", material)
		return
	} else {
		config.CustomResponse(c, 200, "not found id", nil)
		return
	}
}
