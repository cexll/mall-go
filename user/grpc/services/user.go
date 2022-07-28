package services

import (
	"context"
	pb "grpc/protos"
)

type UserService struct {
}

func (t *UserService) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	// 执行数据库操作
	// ...

	resp := pb.AddResponse{
		ErrorCode:    0,
		ErrorMessage: "",
		UserId:       10001,
	}
	return &resp, nil
}
