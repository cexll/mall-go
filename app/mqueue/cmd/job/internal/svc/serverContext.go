package svc

import (
	"github.com/hibiken/asynq"
	"mall-go/app/mqueue/cmd/job/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
	}
}
