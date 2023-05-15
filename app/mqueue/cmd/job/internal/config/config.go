package config

import (
	"github.com/mix-plus/go-mixplus/mrpc"
	"mall-go/common/conf"
)

type Config struct {
	mrpc.ServiceConf
	conf.RedisConf
}
