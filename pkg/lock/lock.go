package lock

import (
	"sync"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"mall-go/common/conf"
)

var (
	once sync.Once
	rs   *redsync.Redsync
)

// NewLock 获取分布式锁
func NewLock(c conf.RedisConf) *redsync.Redsync {
	once.Do(func() {
		client := goredislib.NewClient(&goredislib.Options{
			Addr:        c.Addr,
			Password:    c.Pass,
			DB:          int(c.DataBase),
			DialTimeout: time.Duration(c.Timeout) * time.Second,
		})
		pool := goredis.NewPool(client)
		rs = redsync.New(pool)
	})

	return rs
}
