package svc

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"mall-go/app/community/cmd/rpc/internal/config"
	"mall-go/pkg"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Client
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  pkg.NewRedis(c.RedisConf),
		Db:     pkg.NewGorm(c.DbConf),
	}
}
