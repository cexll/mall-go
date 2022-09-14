package routes

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
)

func Load(router *server.Hertz) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"}, // Allowed domains, need to bring schema
		AllowMethods:     []string{"PUT", "PATCH"},    // Allowed request methods
		AllowHeaders:     []string{"Origin"},          // Allowed request headers
		ExposeHeaders:    []string{"Content-Length"},  // Request headers allowed in the upload_file
		AllowCredentials: true,                        // Whether cookies are attached
		AllowOriginFunc: func(origin string) bool { // Custom domain detection with lower priority than AllowOrigins
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour, // Maximum length of upload_file-side cache preflash requests (seconds)
	}))

	router.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.String(200, "Hello Mall-Go Community!")
	})
	//router.Use(gin.Recovery(), middleware.CorsMiddleware(), middleware.AuthMiddleware())
}
