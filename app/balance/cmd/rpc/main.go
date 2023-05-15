package main

import (
	"flag"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"
	"mall-go/app/balance/cmd/pb"
	"mall-go/app/balance/cmd/rpc/internal/config"
	"mall-go/app/balance/cmd/rpc/internal/server"
	"mall-go/app/balance/cmd/rpc/internal/svc"
	conf "mall-go/common/conf"
	_ "mall-go/pkg/di"
)

var configFile = flag.String("f", "etc/balance.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)
	svc.Context = ctx
	srv := server.NewBalanceServer(ctx)

	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		pb.RegisterBalanceServer(g, srv)
	})

	defer s.Stop()

	s.Start()
}
