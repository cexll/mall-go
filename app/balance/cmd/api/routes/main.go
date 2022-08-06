package routes

import (
	"github.com/gin-gonic/gin"
	"mall-go/app/user/cmd/api/controllers"
	"mall-go/common/middleware"
)

func Load(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.CorsMiddleware())

	user := controllers.UserController{}
	// 注册接口
	router.POST("/user/register", func(ctx *gin.Context) {
		user.Register(ctx)
	})
	// 查看用户详情
	router.POST("/user/getUser", func(ctx *gin.Context) {
		user.GetUser(ctx)
	})
	v1 := router.Group("/user", middleware.AuthMiddleware())
	{
		// 编辑用户信息
		v1.POST("/setUser", func(ctx *gin.Context) {
			user.SetUser(ctx)
		})
		// 注销用户
		v1.POST("/logout", func(ctx *gin.Context) {
			user.Logout(ctx)
		})
	}
}
