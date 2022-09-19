package config

import (
	"github.com/go-ll/mrpc"
	"mall-go/common/conf"
)

type Config struct {
	conf.ApiConf `mapstructure:",squash"`
	conf.JwtAuth
	MerchantsRpcConf mrpc.RpcClientConf
	BalanceRpcConf   mrpc.RpcClientConf
}
