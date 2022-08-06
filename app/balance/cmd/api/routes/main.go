package routes

import (
	"github.com/gin-gonic/gin"
	"mall-go/common/middleware"
)

func Load(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.CorsMiddleware())

}
