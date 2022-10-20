package config

import (
	"github.com/go-ll/mrpc"
	"mall-go/common/conf"
)

type Config struct {
	mrpc.ServiceConf
	conf.RedisConf
}
