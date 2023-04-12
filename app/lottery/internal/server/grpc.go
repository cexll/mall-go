package server

import (
	"mall-go/app/lottery/internal/config"
	"mall-go/app/lottery/internal/service"
	"mall-go/app/lottery/internal/svc"

	lottery "mall-go/app/lottery/api/lottery/v1"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"
)

func NewGrpcServer(c *config.Config, svc *svc.ServiceContext) *mrpc.RpcServer {

	srv := service.NewLotteryService(svc)
	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		// grpc register
		lottery.RegisterLotteryServer(g, srv)
	})

	return s
}
