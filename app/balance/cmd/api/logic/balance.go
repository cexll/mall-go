package logic

import (
	"context"
	"mall-go/app/balance/cmd/api/services"
	"mall-go/app/balance/cmd/pb"
)

type BalanceLogic struct {
	grpcService services.GRPCService
}

type GetBalanceResponse struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Type      int32  `json:"type"`
	Available uint64 `json:"available" default:"0"`
	Frozen    uint64 `json:"frozen" default:"0"`
}

func (t *BalanceLogic) GetBalance(in *pb.GetBalanceRequest) (*GetBalanceResponse, error) {
	cli, err := t.grpcService.NewBalanceRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.GetBalance(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return &GetBalanceResponse{
		Id:        resp.Id,
		UserId:    resp.UserId,
		Type:      resp.Type,
		Available: resp.Available,
		Frozen:    resp.Frozen,
	}, nil
}

func (t *BalanceLogic) GetBalanceChangeList(in *pb.GetBalanceChangeListRequest) (*pb.GetBalanceChangeListResponse, error) {
	cli, err := t.grpcService.NewBalanceRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.GetBalanceChangeList(context.Background(), in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
