package services

import (
	"mall-go/app/balance/cmd/pb"
	"mall-go/common/mrpc"
)

type GRPCService struct {
}

// NewBalanceRpcClient 创建一个BalanceRPC连接
func (t GRPCService) NewBalanceRpcClient() (pb.BalanceClient, error) {
	balance, err := mrpc.NewClient(mrpc.BalanceGRPCAddr, func(options *mrpc.ClientOptions) {})
	if err != nil {
		return nil, err
	}
	return pb.NewBalanceClient(balance.Conn()), nil
}
