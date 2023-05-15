package logic

import (
	"context"

	balancepb "mall-go/app/balance/cmd/pb"
	"mall-go/app/merchants/cmd/api/internal/svc"
	"mall-go/app/merchants/cmd/pb"
)

type MerchantsLogic struct {
}

func (t *MerchantsLogic) JoinMerchant(in *pb.JoinMerchantRequest) (*pb.JoinMerchantResponse, error) {

	resp, err := svc.Context.MerchantsRpc.JoinMerchant(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *MerchantsLogic) UpdateMerchant(in *pb.UpdateMerchantRequest) (*pb.UpdateMerchantResponse, error) {
	resp, err := svc.Context.MerchantsRpc.UpdateMerchant(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *MerchantsLogic) CloseMerchant(in *pb.CloseMerchantRequest) (*pb.CloseMerchantResponse, error) {
	resp, err := svc.Context.MerchantsRpc.CloseMerchant(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *MerchantsLogic) GetMerchant(in *pb.GetMerchantRequest) (*pb.GetMerchantResponse, error) {
	resp, err := svc.Context.MerchantsRpc.GetMerchant(context.Background(), in)
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
	resp, err := svc.Context.BalanceRpc.GetBalance(context.Background(), in)
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
	resp, err := svc.Context.BalanceRpc.GetBalanceChangeList(context.Background(), in)
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
