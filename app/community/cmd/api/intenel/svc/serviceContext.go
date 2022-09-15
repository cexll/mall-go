package svc

import (
	"mall-go/app/community/cmd/api/intenel/config"
	"mall-go/app/community/cmd/pb"
)

var Context *ServiceContext

type ServiceContext struct {
	Config config.Config

	CommunityRpc pb.CommunityClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommunityRpc: nil,
	}
}
