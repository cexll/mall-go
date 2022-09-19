package pkg

import (
	"github.com/go-session/redis"
	"github.com/go-session/session"
	"mall-go/common/conf"
	"time"
)

func NewSession(c conf.RedisConf) *session.Manager {
	opts := redis.Options{
		Addr:        c.Addr,
		Password:    c.Pass,
		DB:          int(c.DataBase),
		DialTimeout: time.Duration(c.Timeout) * time.Second,
	}
	opt := redis.NewRedisStore(&opts)
	return session.NewManager(session.SetStore(opt))
}
