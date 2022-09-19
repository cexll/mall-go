package config

import (
	"github.com/go-ll/mrpc"
	"mall-go/common/conf"
)

type Config struct {
	conf.ApiConf `mapstructure:",squash"`
	conf.JwtAuth
	BalanceRpcConf mrpc.RpcClientConf
}
