package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"mall-go/common/middleware"
	validator2 "mall-go/common/validator"
)

func Load(router *gin.Engine) {

	// register validate
	binding.Validator = new(validator2.DefaultValidator)

	router.Use(gin.Recovery(), middleware.CorsMiddleware(), middleware.AuthMiddleware())
}
