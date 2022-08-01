package services

import (
	"context"
	"errors"
	"mall-go/app/user/pb"
	"mall-go/app/user/rpc/logic"
)

// UserService 用户接口实现
type UserService struct {
	logic logic.UserLogic
}

func (t *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if in.Id == 0 {
		return &pb.GetUserResponse{}, errors.New("id不能为空")
	}
	user, err := t.logic.GetUser(in.Id)
	if err != nil {
		return &pb.GetUserResponse{}, err
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

func (t *UserService) SetUser(ctx context.Context, in *pb.SetUserRequest) (*pb.SetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *UserService) Logout(ctx context.Context, in *pb.LogOutRequest) (*pb.LogOutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *UserService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// 执行数据库操作
	// ...

	resp := pb.RegisterResponse{
		Id: 100,
	}
	return &resp, nil
}
