package server

import (
	"context"
	"mall-go/app/balance/cmd/pb"
	"mall-go/app/balance/cmd/rpc/internal/logic"
	"mall-go/app/balance/cmd/rpc/internal/svc"
	"mall-go/common/errcode"
)

type BalanceServer struct {
	svcCtx *svc.ServiceContext
	logic  logic.BalanceLogic
}

func NewBalanceServer(ctx *svc.ServiceContext) *BalanceServer {
	return &BalanceServer{}
}

func (t *BalanceServer) SubFrozenBalance(ctx context.Context, in *pb.SubFrozenBalanceRequest) (*pb.SubFrozenBalanceResponse, error) {
	status, err := t.logic.SubFrozenBalance(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 102)
	}
	return &pb.SubFrozenBalanceResponse{
		Status: status,
	}, nil
}

func (t *BalanceServer) ReduceFrozenBalance(ctx context.Context, in *pb.ReduceFrozenBalanceRequest) (*pb.ReduceFrozenBalanceResponse, error) {
	status, err := t.logic.ReduceFrozenBalance(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 102)
	}
	return &pb.ReduceFrozenBalanceResponse{
		Status: status,
	}, nil
}

func (t *BalanceServer) SubBalance(ctx context.Context, in *pb.SubBalanceRequest) (*pb.SubBalanceResponse, error) {
	status, err := t.logic.SubBalance(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 102)
	}
	return &pb.SubBalanceResponse{
		Status: status,
	}, nil
}

func (t *BalanceServer) ReduceBalance(ctx context.Context, in *pb.ReduceBalanceRequest) (*pb.ReduceBalanceResponse, error) {
	status, err := t.logic.ReduceBalance(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 102)
	}
	return &pb.ReduceBalanceResponse{
		Status: status,
	}, nil
}

func (t *BalanceServer) GetBalance(ctx context.Context, in *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	balance, err := t.logic.GetBalance(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 102)
	}
	return &pb.GetBalanceResponse{
		Id:        balance.ID,
		UserId:    balance.UserId,
		Type:      int32(balance.Type),
		Available: balance.Available,
		Frozen:    balance.Frozen,
	}, nil
}

func (t *BalanceServer) GetBalanceChangeList(ctx context.Context, in *pb.GetBalanceChangeListRequest) (*pb.GetBalanceChangeListResponse, error) {
	result, err := t.logic.GetBalanceChangeList(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 102)
	}

	return &pb.GetBalanceChangeListResponse{
		TotalCount: result.TotalCount,
		TotalPage:  result.TotalPage,
		List:       result.List,
	}, nil
}
