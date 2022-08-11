package di

import (
	"github.com/go-session/redis"
	"github.com/go-session/session"
	"github.com/mix-go/xdi"
	"mall-go/common/config"
	"time"
)

func init() {
	obj := xdi.Object{
		Name: "session",
		New: func() (i interface{}, e error) {
			opts := redis.Options{
				Addr:        config.Conf.REDIS.RedisAddr,
				Password:    config.Conf.REDIS.RedisPassword,
				DB:          int(config.Conf.REDIS.RedisDatabase),
				DialTimeout: time.Duration(config.Conf.REDIS.RedisTimeOut) * time.Second,
			}
			opt := redis.NewRedisStore(&opts)
			return session.NewManager(session.SetStore(opt)), nil
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}

func Session() (manager *session.Manager) {
	if err := xdi.Populate("session", &manager); err != nil {
		panic(err)
	}
	return
}
