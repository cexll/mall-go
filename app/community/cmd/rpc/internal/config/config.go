package config

import (
	"mall-go/common/conf"

	"github.com/mix-plus/go-mixplus/mrpc"
)

type Config struct {
	mrpc.RpcServerConf `mapstructure:",squash"`
	conf.DbConf
	conf.RedisConf
	CommunityRpcConf mrpc.RpcClientConf
}
