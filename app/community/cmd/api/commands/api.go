package commands

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/mix-go/xcli/flag"
	"github.com/mix-go/xcli/process"
	"mall-go/app/community/cmd/api/routes"
	"mall-go/common/config"
	"mall-go/common/di"
	"time"
)

type APICommand struct {
}

func (t *APICommand) Main() {
	if flag.Match("d", "daemon").Bool() {
		process.Daemon()
	}

	logger := di.Zap()

	addr := config.Conf.API.GinAddr
	mode := config.Conf.API.GinMode
	s := server.Default(
		server.WithHostPorts(addr),
		server.WithExitWaitTime(3*time.Second))

	if mode != "release" {
		s.Use(func(c context.Context, ctx *app.RequestContext) {
			start := time.Now()
			ctx.Next(c)
			end := time.Now()
			latency := end.Sub(start).Microseconds
			logger.Infof("status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s",
				ctx.Response.StatusCode(), latency(), ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host())
		})
	}

	routes.Load(s)

	s.Engine.OnShutdown = append(s.Engine.OnShutdown, func(ctx context.Context) {
		logger.Info("Server shutdown")
	})

	// run
	welcome()
	logger.Infof("Server start at %s", addr)

	s.Spin()
}
