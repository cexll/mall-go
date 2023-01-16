package main

import (
	"flag"
	"mall-go/app/user/cmd/pb"
	"mall-go/app/user/cmd/rpc/internal/config"
	"mall-go/app/user/cmd/rpc/internal/server"
	"mall-go/app/user/cmd/rpc/internal/svc"
	"mall-go/common/conf"

	"github.com/mix-plus/go-mixplus/mrpc"
	"google.golang.org/grpc"

	_ "mall-go/pkg/di"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)
	svc.Context = ctx

	s := mrpc.MustNewServer(c.RpcServerConf, func(g *grpc.Server) {
		pb.RegisterUserServer(g, srv)
	})

	defer s.Stop()

	s.Start()
}
