package services

import (
	"mall-go/app/user/pb"
	"mall-go/common/grpc"
)

type GRPCService struct {
}

// NewUserRpcClient 创建一个user rpc 连接
func (t GRPCService) NewUserRpcClient() (pb.UserClient, error) {
	userConn, err := grpc.NewClient(grpc.UserGRPCAddr, func(options *grpc.ClientOptions) {})
	if err != nil {
		return nil, err
	}
	return pb.NewUserClient(userConn.Conn()), nil
}
