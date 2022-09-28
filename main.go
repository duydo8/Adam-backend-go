package main

import (
	"Adam-backend-go/controller"
	"Adam-backend-go/controller/Admin"
	"Adam-backend-go/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("sign-up", controller.SignUp)
	r.GET("log-in", controller.Login)
	admin := r.Group("admin")
	{
		account := admin.Group("account")
		{
			account.POST("create-account", Admin.CreateAccountAdmin)
			account.GET("verify", Admin.Verify)
			account.GET("find-by-pageable", Admin.AccountAdminFindByPageable)
			account.GET("change-password", Admin.AdminAccountChangePass)
			account.DELETE("delete-by-id", Admin.AdminAccountDeleteById)
		}
		category := admin.Group("category")
		{
			category.POST("create", Admin.AdminCategoryCreate)
			category.GET("find-all", Admin.AdminCategoryFindAll)
			category.GET("find-all-category-parent", Admin.AdminCategoryFindAllCategoryParent)
			category.PUT("update", Admin.AdminCategoryUpdate)
			category.DELETE("deleteById", Admin.AdminCategoryDeleteByID)
			category.GET("find-by-category-parent-id", Admin.AdminCategoryFindCategoryByParentId)
		}
		color := admin.Group("color")
		{
			color.GET("find-all", Admin.ColorAdminFindAll)
			color.POST("create", Admin.ColorAdminCreate)
			color.PUT("update", Admin.ColorAdminUpdate)
			color.DELETE("deleteById", Admin.ColorAdminDelete)
		}
		comment := admin.Group("comment")
		{
			comment.POST("create", Admin.CommentAdminCreate)
		}
		material := admin.Group("material")
		{
			material.GET("find-all", Admin.MaterialAdminFindAll)
			material.POST("create", Admin.MaterialAdminCreate)
			material.PUT("update", Admin.MaterialAdminUpdate)
			material.DELETE("deleteById", Admin.MaterialAdminDelete)
		}
		size := admin.Group("size")
		{
			size.GET("find-all", Admin.SizeAdminFindAll)
			size.POST("create", Admin.SizeAdminCreate)
			size.PUT("update", Admin.SizeAdminUpdate)
			size.DELETE("deleteById", Admin.SizeAdminDelete)
		}
		tag := admin.Group("tag")
		{
			tag.GET("find-all", Admin.TagAdminFindAll)
			tag.POST("create", Admin.TagAdminCreate)
			tag.PUT("update", Admin.TagAdminUpdate)
			tag.DELETE("deleteById", Admin.TagAdminDelete)
		}
		product := admin.Group("product")
		{
			product.POST("create", Admin.ProductAdminCreate)
			product.PUT("update", Admin.ProductAdminUpdate)
			product.DELETE("deleteById", Admin.ProductAdminDelete)
		}

	}
	website := r.Group("/")
	{
		account := website.Group("/account")
		{
			account.POST("create-account", Admin.CreateAccountAdmin)

		}
	}

	r.Run()
}
