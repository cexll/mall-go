package services

import (
	"mall-go/app/balance/cmd/pb"
	"mall-go/common/grpc"
)

type GRPCService struct {
}

// NewUserRpcClient 创建一个user rpc 连接
func (t GRPCService) NewUserRpcClient() (pb.BalanceClient, error) {
	balance, err := grpc.NewClient(grpc.BalanceGRPCAddr, func(options *grpc.ClientOptions) {})
	if err != nil {
		return nil, err
	}
	return pb.NewBalanceClient(balance.Conn()), nil
}
