package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mall-go/app/balance/cmd/api/internal/svc"
	"mall-go/common/middleware"
	validator2 "mall-go/pkg/validator"
)

func Load(router *gin.Engine) {
	// register validate
	binding.Validator = new(validator2.DefaultValidator)
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello")
	})
	router.Use(gin.Recovery(), middleware.CorsMiddleware(), middleware.AuthMiddleware(svc.Context.Jwt.Secret))
}
