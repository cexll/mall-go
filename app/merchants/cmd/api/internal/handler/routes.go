package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mall-go/app/merchants/cmd/api/internal/handler/controllers"
	"mall-go/app/merchants/cmd/api/internal/svc"
	"mall-go/common/middleware"
	validator2 "mall-go/pkg/validator"
)

func Load(router *gin.Engine) {

	// register validate
	binding.Validator = new(validator2.DefaultValidator)

	router.Use(gin.Recovery(), middleware.CorsMiddleware(), middleware.AuthMiddleware(svc.Context.Jwt.Secret))

	merchant := controllers.MerchantsController{}
	r := router.Group("/merchant")
	{
		// 查询商户信息
		r.POST("/getMerchant", func(c *gin.Context) {
			merchant.GetMerchant(c)
		})
		// 加入商户
		r.POST("/join", func(c *gin.Context) {
			merchant.JoinMerchant(c)
		})
		// 更新商户
		r.POST("/update", func(c *gin.Context) {
			merchant.UpdateMerchant(c)
		})
		// 退出商户
		r.POST("/close", func(c *gin.Context) {
			merchant.CloseMerchant(c)
		})
		// 查询余额
		r.POST("/getBalance", func(c *gin.Context) {
			merchant.GetBalance(c)
		})
		// 查询账单
		r.POST("/getBalanceChangeList", func(c *gin.Context) {
			merchant.GetBalanceChangeList(c)
		})
		// 提现余额
		r.POST("/withdrawals", func(c *gin.Context) {
			merchant.Withdrawals(c)
		})
	}

}
