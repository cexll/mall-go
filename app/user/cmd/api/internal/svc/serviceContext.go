package svc

import (
	balancepb "mall-go/app/balance/cmd/pb"
	"mall-go/app/user/cmd/api/internal/config"
	"mall-go/app/user/cmd/pb"
	"mall-go/pkg/jwtx"

	"github.com/mix-plus/go-mixplus/mrpc"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config
	Jwt    *jwtx.Jwt

	UserRpc    pb.UserClient
	BalanceRpc balancepb.BalanceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Jwt:        jwtx.NewJwt(c.JwtAuth),
		UserRpc:    pb.NewUserClient(mrpc.MustNewClient(c.UserRpcConf).Conn()),
		BalanceRpc: balancepb.NewBalanceClient(mrpc.MustNewClient(c.BalanceRpcConf).Conn()),
	}
}
