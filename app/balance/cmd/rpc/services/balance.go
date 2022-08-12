package services

import (
	"context"
	"mall-go/app/balance/cmd/pb"
	"mall-go/app/balance/cmd/rpc/logic"
	"mall-go/common/grpc"
)

type BalanceService struct {
	logic logic.BalanceLogic
}

func (t *BalanceService) SubFrozenBalance(ctx context.Context, in *pb.SubFrozenBalanceRequest) (*pb.SubFrozenBalanceResponse, error) {
	status, err := t.logic.SubFrozenBalance(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 102)
	}
	return &pb.SubFrozenBalanceResponse{
		Status: status,
	}, nil
}

func (t *BalanceService) ReduceFrozenBalance(ctx context.Context, in *pb.ReduceFrozenBalanceRequest) (*pb.ReduceFrozenBalanceResponse, error) {
	status, err := t.logic.ReduceFrozenBalance(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 102)
	}
	return &pb.ReduceFrozenBalanceResponse{
		Status: status,
	}, nil
}

func (t *BalanceService) SubBalance(ctx context.Context, in *pb.SubBalanceRequest) (*pb.SubBalanceResponse, error) {
	status, err := t.logic.SubBalance(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 102)
	}
	return &pb.SubBalanceResponse{
		Status: status,
	}, nil
}

func (t *BalanceService) ReduceBalance(ctx context.Context, in *pb.ReduceBalanceRequest) (*pb.ReduceBalanceResponse, error) {
	status, err := t.logic.ReduceBalance(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 102)
	}
	return &pb.ReduceBalanceResponse{
		Status: status,
	}, nil
}

func (t *BalanceService) GetBalance(ctx context.Context, in *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	balance, err := t.logic.GetBalance(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 102)
	}
	return &pb.GetBalanceResponse{
		Id:        balance.ID,
		UserId:    balance.UserId,
		Type:      int32(balance.Type),
		Available: balance.Available,
		Frozen:    balance.Frozen,
	}, nil
}
