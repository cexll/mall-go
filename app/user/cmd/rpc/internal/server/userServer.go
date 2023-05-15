package server

import (
	"context"
	"errors"

	"mall-go/app/user/cmd/pb"
	"mall-go/app/user/cmd/rpc/internal/logic"
	"mall-go/app/user/cmd/rpc/internal/svc"
	"mall-go/common/errcode"
)

// UserServer 用户接口实现
type UserServer struct {
	svcCtx *svc.ServiceContext
	logic  logic.UserLogic
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (t *UserServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if in.Id == 0 {
		return nil, errcode.GrpcError(errors.New("id不能为空"), 100)
	}
	user, err := t.logic.GetUser(in.Id)
	if err != nil {
		return nil, errcode.GrpcError(err, 100)
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

func (t *UserServer) SetUser(ctx context.Context, in *pb.SetUserRequest) (*pb.SetUserResponse, error) {
	user, err := t.logic.SetUser(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 100)
	}

	return &pb.SetUserResponse{
		Status: user,
	}, nil
}

func (t *UserServer) Logout(ctx context.Context, in *pb.LogOutRequest) (*pb.LogOutResponse, error) {
	logout, err := t.logic.Logout(in.Id)
	if err != nil {
		return nil, errcode.GrpcError(err, 100)
	}

	return &pb.LogOutResponse{
		Status: logout,
	}, nil
}

func (t *UserServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := t.logic.Register(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 100)
	}

	return &pb.RegisterResponse{
		Id: id,
	}, nil
}

func (t *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := t.logic.Login(in)
	if err != nil {
		return nil, errcode.GrpcError(err, 100)
	}
	return &pb.LoginResponse{
		Id:        user.ID,
		Nickname:  user.Nickname,
		AvatarUrl: user.AvatarUrl,
		Mobile:    user.Mobile,
	}, nil
}
