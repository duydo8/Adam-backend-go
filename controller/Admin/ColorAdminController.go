package Admin

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"github.com/gin-gonic/gin"
	"time"
)

func ColorAdminCreate(c *gin.Context) {
	var colorCreate model.ColorCreate
	var color model.Color
	c.ShouldBindJSON(&colorCreate)
	initializers.DB.Where("color_name=?", colorCreate.ColorName).First(&color)
	if color.ID != 0 {
		config.CustomResponse(c, 200, "duplicate color name", nil)
		return
	}
	color.CreateDate = time.Now()
	color.ColorName = colorCreate.ColorName
	color.Status = true
	initializers.DB.Save(&color)
	config.CustomResponse(c, 200, "success", color)
	return

}

func ColorAdminUpdate(c *gin.Context) {
	var color model.Color
	var color2 model.Color
	var color1 model.ColorCreate
	c.ShouldBindJSON(&color1)
	initializers.DB.Where("id=?", color1.ID).First(&color)
	if color1.ID == 0 {
		config.CustomResponse(c, 200, "not found", nil)
		return
	}
	initializers.DB.Where("color_name=?", color1.ColorName).First(&color2)
	if color2.ID != 0 {
		config.CustomResponse(c, 200, "duplicate color name", nil)
		return
	}
	color.ColorName = color1.ColorName

	initializers.DB.Save(&color)
	config.CustomResponse(c, 200, "success", color)
	return
}
func ColorAdminFindAll(c *gin.Context) {
	var color []model.Color
	initializers.DB.Where("status =1").Find(&color)
	config.CustomResponse(c, 200, "success", color)
	return

}
func ColorAdminDelete(c *gin.Context) {
	var color model.Color
	id := c.Query("id")
	initializers.DB.First(&model.Color{}, id).Scan(&color)
	if color.ID != 0 {
		color.Status = false
		initializers.DB.Save(&color)
		config.CustomResponse(c, 200, "success", color)
		return
	} else {
		config.CustomResponse(c, 200, "not found id", nil)
		return
	}
}
