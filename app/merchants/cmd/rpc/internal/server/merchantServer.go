package server

import (
	"context"
	"mall-go/app/merchants/cmd/pb"
	"mall-go/app/merchants/cmd/rpc/internal/logic"
	"mall-go/app/merchants/cmd/rpc/internal/svc"
	"mall-go/common/errcode"
)

type MerchantsServer struct {
	svcCtx *svc.ServiceContext
	logic  logic.MerchantsLogic
}

func NewMerchantsServer(ctx *svc.ServiceContext) *MerchantsServer {
	return &MerchantsServer{
		svcCtx: ctx,
	}
}

func (m *MerchantsServer) GetMerchant(c context.Context, in *pb.GetMerchantRequest) (*pb.GetMerchantResponse, error) {
	merchant, err := m.logic.GetMerchant(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 201)
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

func (m *MerchantsServer) JoinMerchant(c context.Context, in *pb.JoinMerchantRequest) (*pb.JoinMerchantResponse, error) {
	status, err := m.logic.JoinMerchant(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 201)
	}
	return &pb.JoinMerchantResponse{
		Status: status,
	}, nil
}

func (m *MerchantsServer) UpdateMerchant(c context.Context, in *pb.UpdateMerchantRequest) (*pb.UpdateMerchantResponse, error) {
	status, err := m.logic.UpdateMerchant(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 201)
	}

	return &pb.UpdateMerchantResponse{
		Status: status,
	}, nil
}

func (m *MerchantsServer) CloseMerchant(c context.Context, in *pb.CloseMerchantRequest) (*pb.CloseMerchantResponse, error) {
	status, err := m.logic.CloseMerchant(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 201)
	}

	return &pb.CloseMerchantResponse{
		Status: status,
	}, nil
}
