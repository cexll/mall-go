package svc

import (
	"github.com/mix-plus/go-mixplus/mrpc"
	"mall-go/app/balance/cmd/api/internal/config"
	"mall-go/app/balance/cmd/pb"
	"mall-go/pkg/jwtx"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config
	Jwt    *jwtx.Jwt

	BalanceRpc pb.BalanceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Jwt:        jwtx.NewJwt(c.JwtAuth),
		BalanceRpc: pb.NewBalanceClient(mrpc.MustNewClient(c.BalanceRpcConf).Conn()),
	}
}
