package lock

import (
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"mall-go/common/config"
	"sync"
)

var (
	once sync.Once
	rs   *redsync.Redsync
)

// NewLock 获取分布式锁
func NewLock() *redsync.Redsync {
	once.Do(func() {
		client := goredislib.NewClient(&goredislib.Options{
			Addr:     config.Conf.REDIS.RedisAddr,
			Password: config.Conf.REDIS.RedisPassword,
			DB:       int(config.Conf.REDIS.RedisDatabase),
		})
		pool := goredis.NewPool(client)
		rs = redsync.New(pool)
	})

	return rs
}
