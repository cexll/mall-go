package logic

import (
	"context"
	balancepb "mall-go/app/balance/cmd/pb"
	"mall-go/app/merchants/cmd/api/services"
	"mall-go/app/merchants/cmd/pb"
)

type MerchantsLogic struct {
	grpcService services.GRPCService
}

func (t *MerchantsLogic) JoinMerchant(in *pb.JoinMerchantRequest) (*pb.JoinMerchantResponse, error) {
	cli, err := t.grpcService.NewMerchantsRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.JoinMerchant(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *MerchantsLogic) UpdateMerchant(in *pb.UpdateMerchantRequest) (*pb.UpdateMerchantResponse, error) {
	cli, err := t.grpcService.NewMerchantsRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.UpdateMerchant(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *MerchantsLogic) CloseMerchant(in *pb.CloseMerchantRequest) (*pb.CloseMerchantResponse, error) {
	cli, err := t.grpcService.NewMerchantsRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.CloseMerchant(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *MerchantsLogic) GetMerchant(in *pb.GetMerchantRequest) (*pb.GetMerchantResponse, error) {
	cli, err := t.grpcService.NewMerchantsRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.GetMerchant(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type GetBalanceResponse struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Type      int32  `json:"type"`
	Available uint64 `json:"available" default:"0"`
	Frozen    uint64 `json:"frozen" default:"0"`
}

func (t *MerchantsLogic) GetBalance(in *balancepb.GetBalanceRequest) (*GetBalanceResponse, error) {
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

func (t *MerchantsLogic) GetBalanceChangeList(in *balancepb.GetBalanceChangeListRequest) (*balancepb.GetBalanceChangeListResponse, error) {
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

type WithdrawalsRequest struct{}

type WithdrawalsResponse struct{}

func (t *MerchantsLogic) Withdrawals(in *WithdrawalsRequest) (*WithdrawalsResponse, error) {
	// 申请提现
	// 扣除余额
	// 同意提现

	// 拒绝提现
	// 返还余额
	return nil, nil
}
