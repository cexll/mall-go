package server

import (
	"net/http"

	"mall-go/app/lottery/internal/service"
	"mall-go/app/lottery/internal/svc"

	"github.com/google/wire"
	"github.com/mix-plus/go-mixplus/mrpc"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

type AppServer struct {
	SvcCtx     *svc.ServiceContext
	HttpServer *http.Server
	GrpcServer *mrpc.RpcServer

	HelloService *service.LotteryService
}

func NewApp(svcCtx *svc.ServiceContext, helloService *service.LotteryService, hs *http.Server, gs *mrpc.RpcServer) (*AppServer, error) {
	return &AppServer{
		SvcCtx:       svcCtx,
		HelloService: helloService,
		HttpServer:   hs,
		GrpcServer:   gs,
	}, nil
}

func (a *AppServer) Run() {

	go func() {
		a.HttpServer.ListenAndServe()
	}()

	a.GrpcServer.Start()

	defer a.GrpcServer.Stop()
}
