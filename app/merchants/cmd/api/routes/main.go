package routes

import (
	"github.com/gin-gonic/gin"
	"mall-go/common/middleware"
)

func Load(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.CorsMiddleware())
	// 加入商户
	// 更新商户
	// 退出商户
	// 查询余额
	// 提现余额
	// 查询账单

}
