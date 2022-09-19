package config

import (
	"github.com/go-ll/mrpc"
	"mall-go/common/conf"
)

type Config struct {
	mrpc.RpcServerConf `mapstructure:",squash"`
	conf.DbConf
	conf.RedisConf
	CommunityRpcConf mrpc.RpcClientConf
}
