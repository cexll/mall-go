package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mall-go/app/balance/cmd/api/controllers"
	"mall-go/common/middleware"
	validator2 "mall-go/common/validator"
)

func Load(router *gin.Engine) {

	// register validate
	binding.Validator = new(validator2.DefaultValidator)

	router.Use(gin.Recovery(), middleware.CorsMiddleware(), middleware.AuthMiddleware())

	balance := controllers.BalanceController{}
	// 查询余额
	router.POST("/balance/getBalance", func(c *gin.Context) {
		balance.GetBalance(c)
	})
	// 查询余额变动账单
	router.POST("/balance/getBalanceChangeList", func(c *gin.Context) {
		balance.GetBalanceChangeList(c)
	})
	// 充值
	router.POST("/balance/recharge", func(c *gin.Context) {
		balance.Recharge(c)
	})
	// 提现
	router.POST("/balance/withdraw", func(c *gin.Context) {
		balance.Withdraw(c)
	})
}
