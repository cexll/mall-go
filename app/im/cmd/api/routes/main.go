package routes

import (
	"github.com/gin-gonic/gin"
	"mall-go/app/im/cmd/api/controllers"
	"mall-go/common/middleware"
)

func Load(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.CorsMiddleware())

	router.GET("websocket",
		func(ctx *gin.Context) {
			ws := controllers.WebSocketController{}
			ws.Index(ctx)
		},
	)
}
