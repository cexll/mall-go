package svc

import (
	balancepb "mall-go/app/balance/cmd/pb"
	"mall-go/app/merchants/cmd/api/internal/config"
	"mall-go/app/merchants/cmd/pb"
	"mall-go/pkg/jwtx"

	"github.com/mix-plus/go-mixplus/mrpc"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config
	Jwt    *jwtx.Jwt

	MerchantsRpc pb.MerchantsClient
	BalanceRpc   balancepb.BalanceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Jwt:          jwtx.NewJwt(c.JwtAuth),
		MerchantsRpc: pb.NewMerchantsClient(mrpc.MustNewClient(c.MerchantsRpcConf).Conn()),
		BalanceRpc:   balancepb.NewBalanceClient(mrpc.MustNewClient(c.BalanceRpcConf).Conn()),
	}
}
