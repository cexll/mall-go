package commands

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mix-go/xcli/flag"
	"github.com/mix-go/xcli/process"
	"mall-go/app/order/cmd/api/routes"
	"mall-go/common/config"
	"mall-go/common/di"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type APICommand struct {
}

func (t *APICommand) Main() {
	if flag.Match("d", "daemon").Bool() {
		process.Daemon()
	}

	logger := di.Zap()
	server := di.Server()
	addr := config.Conf.API.GinAddr
	mode := config.Conf.API.GinMode

	// server
	gin.SetMode(mode)
	router := gin.New()
	if mode != gin.ReleaseMode {
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
	routes.Load(router)

	server.Addr = addr
	server.Handler = router

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

	// run
	welcome()
	logger.Infof("Server start at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && !strings.Contains(err.Error(), "http: Server closed") {
		panic(err)
	}
}
