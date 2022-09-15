package services

import (
	balancepb "mall-go/app/balance/cmd/pb"
	"mall-go/app/merchants/cmd/pb"
	"mall-go/common/mrpc"
)

type GRPCService struct {
}

// NewBalanceRpcClient 创建一个BalanceRPC 连接
func (t GRPCService) NewBalanceRpcClient() (balancepb.BalanceClient, error) {
	balance, err := mrpc.NewClient(mrpc.BalanceGRPCAddr, func(options *mrpc.ClientOptions) {})
	if err != nil {
		return nil, err
	}
	return balancepb.NewBalanceClient(balance.Conn()), nil
}

// NewMerchantsRpcClient 创建一个 MerchantsRPC 连接
func (t GRPCService) NewMerchantsRpcClient() (pb.MerchantsClient, error) {
	merchants, err := mrpc.NewClient(mrpc.MerchantsGRPCAddr, func(options *mrpc.ClientOptions) {})
	if err != nil {
		return nil, err
	}
	return pb.NewMerchantsClient(merchants.Conn()), nil
}
