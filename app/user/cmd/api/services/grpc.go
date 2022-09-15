package services

import (
	balancepb "mall-go/app/balance/cmd/pb"
	"mall-go/app/user/cmd/pb"
	"mall-go/common/mrpc"
)

type GRPCService struct {
}

// NewUserRpcClient 创建一个 UserRPC 连接
func (t GRPCService) NewUserRpcClient() (pb.UserClient, error) {
	userConn, err := mrpc.NewClient(mrpc.UserGRPCAddr, func(options *mrpc.ClientOptions) {})
	if err != nil {
		return nil, err
	}
	return pb.NewUserClient(userConn.Conn()), nil
}

// NewBalanceRpcClient 创建一个 BalanceRPC 连接
func (t GRPCService) NewBalanceRpcClient() (balancepb.BalanceClient, error) {
	balance, err := mrpc.NewClient(mrpc.BalanceGRPCAddr, func(options *mrpc.ClientOptions) {})
	if err != nil {
		return nil, err
	}
	return balancepb.NewBalanceClient(balance.Conn()), nil
}
