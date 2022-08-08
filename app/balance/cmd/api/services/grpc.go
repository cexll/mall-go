package services

import (
	"mall-go/app/balance/cmd/pb"
	"mall-go/common/grpc"
)

type GRPCService struct {
}

// NewBalanceRpcClient 创建一个BalanceRPC连接
func (t GRPCService) NewBalanceRpcClient() (pb.BalanceClient, error) {
	balance, err := grpc.NewClient(grpc.BalanceGRPCAddr, func(options *grpc.ClientOptions) {})
	if err != nil {
		return nil, err
	}
	return pb.NewBalanceClient(balance.Conn()), nil
}
