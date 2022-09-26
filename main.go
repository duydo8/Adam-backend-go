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
		admin.POST("create-account", Admin.CreateAccountAdmin)
		admin.GET("verify", Admin.Verify)
		admin.GET("find-by-pageable", Admin.AccountAdminFindByPageable)
		admin.GET("change-password", Admin.AdminAccountChangePass)
		admin.DELETE("delete-by-id", Admin.AdminAccountDeleteById)
	}

	r.Run()
}
