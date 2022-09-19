package svc

import (
	"github.com/go-redis/redis"
	"github.com/go-redsync/redsync/v4"
	"github.com/mix-go/xsql"
	"gorm.io/gorm"
	"mall-go/app/balance/cmd/rpc/internal/config"
	"mall-go/pkg"
	"mall-go/pkg/lock"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config

	Redis   *redis.Client
	Db      *gorm.DB
	Xsql    *xsql.DB
	RedSync *redsync.Redsync
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Redis:   pkg.NewRedis(c.RedisConf),
		Db:      pkg.NewGorm(c.DbConf),
		Xsql:    pkg.NewXsql(c.DbConf),
		RedSync: lock.NewLock(c.RedisConf),
	}
}
