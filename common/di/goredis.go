package di

import (
	"github.com/go-redis/redis"
	"github.com/mix-go/xdi"
	"mall-go/common/config"
	"time"
)

func init() {
	obj := xdi.Object{
		Name: "goredis",
		New: func() (i interface{}, e error) {
			opt := redis.Options{
				Addr:        config.Conf.REDIS.RedisAddr,
				Password:    config.Conf.REDIS.RedisPassword,
				DB:          int(config.Conf.REDIS.RedisDatabase),
				DialTimeout: time.Duration(config.Conf.REDIS.RedisTimeOut) * time.Second,
			}
			return redis.NewClient(&opt), nil
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}

func GoRedis() (client *redis.Client) {
	if err := xdi.Populate("goredis", &client); err != nil {
		panic(err)
	}
	return
}
