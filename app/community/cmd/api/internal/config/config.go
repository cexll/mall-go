package config

import (
	"mall-go/common/conf"

	"github.com/mix-plus/go-mixplus/mrpc"
)

type Config struct {
	conf.ApiConf `mapstructure:",squash"`
	conf.JwtAuth
	CommunityRpcConf mrpc.RpcClientConf
}
