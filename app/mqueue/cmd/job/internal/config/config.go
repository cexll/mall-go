package config

import (
	"mall-go/common/conf"

	"github.com/mix-plus/go-mixplus/mrpc"
)

type Config struct {
	mrpc.ServiceConf
	conf.RedisConf
}
