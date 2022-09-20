package svc

import (
	"github.com/go-redis/redis"
	"github.com/mix-go/xsql"
	"gorm.io/gorm"
	"mall-go/app/user/cmd/rpc/internal/config"
	"mall-go/pkg"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Client
	Db     *gorm.DB
	Xsql   *xsql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  pkg.NewRedis(c.RedisConf),
		Db:     pkg.NewGorm(c.DbConf),
		Xsql:   pkg.NewXsql(c.DbConf),
	}
}
