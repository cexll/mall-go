package logic

import (
	"context"
	"mall-go/app/user/cmd/api/services"
	"mall-go/app/user/cmd/pb"
)

type UserLogic struct {
	grpcService services.GRPCService
}

func (t UserLogic) Register(req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	cli, err := t.grpcService.NewUserRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.Register(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
