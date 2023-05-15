package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"mall-go/app/balance/cmd/api/internal/config"
	"mall-go/app/balance/cmd/api/internal/handler"
	"mall-go/app/balance/cmd/api/internal/svc"
	conf "mall-go/common/conf"
	"mall-go/pkg/di"
	_ "mall-go/pkg/di"
)

var configFile = flag.String("f", "etc/balance.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	logger := di.Zap()
	server := di.Server()

	svc.Context = svc.NewServiceContext(c)

	gin.SetMode(c.Mode)

	router := gin.New()
	if c.Mode != gin.ReleaseMode {
		handlerFunc := gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: func(params gin.LogFormatterParams) string {
				return fmt.Sprintf("%s|%s|%d|%s\n",
					params.Method,
					params.Path,
					params.StatusCode,
					params.ClientIP,
				)
			},
			Output: &di.ZapOutput{Logger: logger},
		})
		router.Use(handlerFunc)
	}

	server.Addr = c.Addr
	server.Handler = router

	handler.Load(router)

	// signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		logger.Info("Server shutdown")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(ctx); err != nil {
			logger.Errorf("Server shutdown error: %s", err)
		}
	}()

	logger.Infof("Server start at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && !strings.Contains(err.Error(), "http: Server closed") {
		panic(err)
	}
}
