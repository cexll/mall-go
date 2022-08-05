package services

import (
	"context"
	"mall-go/app/balance/cmd/pb"
)

type BalanceService struct {
}

func (t *BalanceService) SubBalance(ctx context.Context, request *pb.SubBalanceRequest) (*pb.SubBalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *BalanceService) ReduceBalance(ctx context.Context, request *pb.ReduceBalanceRequest) (*pb.ReduceBalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *BalanceService) GetBalance(ctx context.Context, in *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {

	return &pb.GetBalanceResponse{
		Id:        1,
		UserId:    2,
		Type:      2,
		Available: 100.00,
		Freeze:    0.00,
	}, nil
}
