package pkg

import (
	"github.com/go-redis/redis"
	"mall-go/common/conf"
	"time"
)

func NewRedis(c conf.RedisConf) (client *redis.Client) {
	return redis.NewClient(&redis.Options{
		Addr:        c.Addr,
		Password:    c.Pass,
		DB:          int(c.DataBase),
		DialTimeout: time.Duration(c.Timeout) * time.Second,
	})
}
