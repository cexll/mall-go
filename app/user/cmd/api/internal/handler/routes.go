package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mall-go/app/user/cmd/api/internal/handler/controllers"
	"mall-go/app/user/cmd/api/internal/svc"
	"mall-go/common/middleware"
	validator2 "mall-go/pkg/validator"
)

func Load(router *gin.Engine) {
	// register validate
	binding.Validator = new(validator2.DefaultValidator)
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})
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
	// 登陆
	router.POST("/user/login", func(ctx *gin.Context) {
		user.Login(ctx)
	})
	v1 := router.Group("/user", middleware.AuthMiddleware(svc.Context.Jwt.Secret))
	{
		// 编辑用户信息
		v1.POST("/setUser", func(ctx *gin.Context) {
			user.SetUser(ctx)
		})
		// 注销用户
		v1.POST("/logout", func(ctx *gin.Context) {
			user.Logout(ctx)
		})
		// 查询余额
		v1.POST("/getBalance", func(ctx *gin.Context) {
			user.GetBalance(ctx)
		})
		// 查询余额变动记录
		v1.POST("/getBalanceChangeList", func(ctx *gin.Context) {
			user.GetBalanceChangeList(ctx)
		})
		// 充值
		v1.POST("/recharge", func(ctx *gin.Context) {
			user.Recharge(ctx)
		})
	}
}
