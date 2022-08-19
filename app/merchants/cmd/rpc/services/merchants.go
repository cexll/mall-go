package services

import (
	"context"
	"mall-go/app/merchants/cmd/pb"
	"mall-go/app/merchants/cmd/rpc/logic"
	"mall-go/common/grpc"
)

type MerchantsService struct {
	logic logic.MerchantsLogic
}

func (m *MerchantsService) JoinMerchant(c context.Context, in *pb.JoinMerchantRequest) (*pb.JoinMerchantResponse, error) {
	status, err := m.logic.JoinMerchant(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 201)
	}
	return &pb.JoinMerchantResponse{
		Status: status,
	}, nil
}

func (m *MerchantsService) UpdateMerchant(c context.Context, in *pb.UpdateMerchantRequest) (*pb.UpdateMerchantResponse, error) {
	status, err := m.logic.UpdateMerchant(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 201)
	}

	return &pb.UpdateMerchantResponse{
		Status: status,
	}, nil
}

func (m *MerchantsService) CloseMerchant(c context.Context, in *pb.CloseMerchantRequest) (*pb.CloseMerchantResponse, error) {
	status, err := m.logic.CloseMerchant(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 201)
	}

	return &pb.CloseMerchantResponse{
		Status: status,
	}, nil
}
