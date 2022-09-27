package Admin

import (
	"Adam-backend-go/config"
	"Adam-backend-go/initializers"
	"Adam-backend-go/model"
	"github.com/gin-gonic/gin"
	"time"
)

func TagAdminFindAll(c *gin.Context) {
	var tag []model.Tag
	initializers.DB.Where("status =1").Find(&tag)
	config.CustomResponse(c, 200, "success", tag)
	return

}
func TagAdminCreate(c *gin.Context) {
	var tag model.TagCreate
	var tag1 model.Tag
	c.ShouldBindJSON(&tag)
	initializers.DB.Where("tag_name=?", tag.TagName).First(&tag1)
	if tag1.ID != 0 {
		config.CustomResponse(c, 200, "duplicate tag name", nil)
		return
	}
	tag1.CreateDate = time.Now()
	tag1.TagName = tag.TagName
	tag1.Status = true
	initializers.DB.Save(&tag1)
	config.CustomResponse(c, 200, "success", tag1)
	return

}

func TagAdminUpdate(c *gin.Context) {
	var tag model.Tag
	var tag2 model.Tag
	var tag1 model.TagCreate
	c.ShouldBindJSON(&tag1)
	initializers.DB.Where("id=?", tag1.ID).First(&tag)
	if tag1.ID == 0 {
		config.CustomResponse(c, 200, "not found", nil)
		return
	}
	initializers.DB.Where("tag_name=?", tag1.TagName).First(&tag2)
	if tag2.ID != 0 {
		config.CustomResponse(c, 200, "duplicate tag name", nil)
		return
	}
	tag.TagName = tag1.TagName

	initializers.DB.Save(&tag)
	config.CustomResponse(c, 200, "success", tag)
	return
}
func TagAdminDelete(c *gin.Context) {
	var tag model.Tag
	id := c.Query("id")
	initializers.DB.First(&model.Tag{}, id).Scan(&tag)
	if tag.ID != 0 {
		tag.Status = false
		initializers.DB.Save(&tag)
		config.CustomResponse(c, 200, "success", tag)
		return
	} else {
		config.CustomResponse(c, 200, "not found id", nil)
		return
	}
}
