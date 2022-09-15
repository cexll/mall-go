package services

import (
	"context"
	"mall-go/app/merchants/cmd/pb"
	"mall-go/app/merchants/cmd/rpc/logic"
	"mall-go/common/mrpc"
)

type MerchantsService struct {
	logic logic.MerchantsLogic
}

func (m *MerchantsService) GetMerchant(c context.Context, in *pb.GetMerchantRequest) (*pb.GetMerchantResponse, error) {
	merchant, err := m.logic.GetMerchant(in)
	if err != nil {
		return nil, mrpc.GrpcError(err, 201)
	}
	return &pb.GetMerchantResponse{
		Id:       merchant.ID,
		UserId:   merchant.UserId,
		ShopName: merchant.ShopName,
		ShopLogo: merchant.ShopLogo,
		Mobile:   merchant.Mobile,
		Address:  merchant.Address,
		Remark:   merchant.Remark,
		Status:   int64(merchant.Status),
	}, nil
}

func (m *MerchantsService) JoinMerchant(c context.Context, in *pb.JoinMerchantRequest) (*pb.JoinMerchantResponse, error) {
	status, err := m.logic.JoinMerchant(in)
	if err != nil {
		return nil, mrpc.GrpcError(err, 201)
	}
	return &pb.JoinMerchantResponse{
		Status: status,
	}, nil
}

func (m *MerchantsService) UpdateMerchant(c context.Context, in *pb.UpdateMerchantRequest) (*pb.UpdateMerchantResponse, error) {
	status, err := m.logic.UpdateMerchant(in)
	if err != nil {
		return nil, mrpc.GrpcError(err, 201)
	}

	return &pb.UpdateMerchantResponse{
		Status: status,
	}, nil
}

func (m *MerchantsService) CloseMerchant(c context.Context, in *pb.CloseMerchantRequest) (*pb.CloseMerchantResponse, error) {
	status, err := m.logic.CloseMerchant(in)
	if err != nil {
		return nil, mrpc.GrpcError(err, 201)
	}

	return &pb.CloseMerchantResponse{
		Status: status,
	}, nil
}
