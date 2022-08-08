package services

import (
	"context"
	"errors"
	"mall-go/app/user/cmd/pb"
	"mall-go/app/user/cmd/rpc/logic"
	"mall-go/common/grpc"
)

// UserService 用户接口实现
type UserService struct {
	logic logic.UserLogic
}

func (t UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if in.Id == 0 {
		return nil, grpc.GrpcError(errors.New("id不能为空"), 100)
	}
	user, err := t.logic.GetUser(in.Id)
	if err != nil {
		return nil, grpc.GrpcError(err, 100)
	}
	return &pb.GetUserResponse{
		Id:        user.ID,
		Nickname:  user.Nickname,
		AvatarUrl: user.AvatarUrl,
		Mobile:    user.Mobile,
		Signature: user.Signature,
		Status:    int64(user.Status),
		IsDelete:  int64(user.IsDelete),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (t UserService) SetUser(ctx context.Context, in *pb.SetUserRequest) (*pb.SetUserResponse, error) {
	bool, err := t.logic.SetUser(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 100)
	}

	return &pb.SetUserResponse{
		Status: bool,
	}, nil
}

func (t UserService) Logout(ctx context.Context, in *pb.LogOutRequest) (*pb.LogOutResponse, error) {
	bool, err := t.logic.Logout(in.Id)
	if err != nil {
		return nil, grpc.GrpcError(err, 100)
	}

	return &pb.LogOutResponse{
		Status: bool,
	}, nil
}

func (t UserService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := t.logic.Register(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 100)
	}

	return &pb.RegisterResponse{
		Id: id,
	}, nil
}

func (t UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := t.logic.Login(in)
	if err != nil {
		return nil, grpc.GrpcError(err, 100)
	}
	return &pb.LoginResponse{
		Id:        user.ID,
		Nickname:  user.Nickname,
		AvatarUrl: user.AvatarUrl,
		Mobile:    user.Mobile,
	}, nil
}
