package routes

import (
	"github.com/cexll/mall-go/controllers"
	"github.com/cexll/mall-go/middleware"
	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) {
	router.Use(gin.Recovery()) // error handle

	router.POST("hello", middleware.CorsMiddleware(), func(ctx *gin.Context) {
		hello := controllers.HelloController{}
		hello.Index(ctx)
	})

	// 用户
	user := router.Group("user", middleware.CorsMiddleware())
	{
		userController := controllers.UserController{}
		// 获取
		user.GET("get", func(ctx *gin.Context) {
			userController.Find(ctx)
		})
		// 创建
		user.POST("create", func(ctx *gin.Context) {
			userController.Create(ctx)
		})
	}
}
