package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"mall-go/app/community/cmd/api/internal/config"
	"mall-go/app/community/cmd/api/internal/handler"
	"mall-go/app/community/cmd/api/internal/svc"
	conf "mall-go/common/conf"
	_ "mall-go/pkg/di"
	"time"
)

var configFile = flag.String("f", "etc/community.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	svc.Context = svc.NewServiceContext(c)

	s := server.Default(
		server.WithHostPorts(c.Addr),
		server.WithExitWaitTime(3*time.Second))

	if c.Mode != "release" {
		s.Use(func(c context.Context, ctx *app.RequestContext) {
			start := time.Now()
			ctx.Next(c)
			end := time.Now()
			latency := end.Sub(start).Microseconds
			logger.Infof(fmt.Sprintf("status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s", ctx.Response.StatusCode(), latency(), ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host()))
		})
	}

	handler.Load(s)

	s.Engine.OnShutdown = append(s.Engine.OnShutdown, func(ctx context.Context) {
		logger.Info("Server shutdown")
	})

	s.Spin()
}
