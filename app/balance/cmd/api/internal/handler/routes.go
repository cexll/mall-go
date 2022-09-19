package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mall-go/app/balance/cmd/api/internal/svc"
	validator2 "mall-go/pkg/validator"

	"mall-go/common/middleware"
)

func Load(router *gin.Engine) {

	// register validate
	binding.Validator = new(validator2.DefaultValidator)

	router.Use(gin.Recovery(), middleware.CorsMiddleware(), middleware.AuthMiddleware(svc.Context.Jwt.Secret))
}
