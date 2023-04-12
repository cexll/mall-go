package server

import (
	"context"
	"net/http"

	lottery "mall-go/app/lottery/api/lottery/v1"
	"mall-go/app/lottery/internal/config"
	"mall-go/app/lottery/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewHttpServer(c *config.Config, srv *service.LotteryService) *http.Server {
	httpServer := &http.Server{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	err := lottery.RegisterLotteryHandlerServer(ctx, mux, srv)
	if err != nil {
		panic(err)
	}
	httpServer.Addr = c.ApiConf.Addr
	httpServer.Handler = mux

	return httpServer
}
