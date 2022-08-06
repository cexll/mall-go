package logic

import (
	"context"
	"mall-go/app/user/cmd/api/services"
	"mall-go/app/user/cmd/pb"
	"mall-go/common/jwtx"
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

func (t UserLogic) GetUser(req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	cli, err := t.grpcService.NewUserRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.GetUser(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t UserLogic) SetUser(req *pb.SetUserRequest) (*pb.SetUserResponse, error) {
	cli, err := t.grpcService.NewUserRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.SetUser(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t UserLogic) Logout(req *pb.LogOutRequest) (*pb.LogOutResponse, error) {
	cli, err := t.grpcService.NewUserRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.Logout(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type LoginResponse struct {
	Jwt       string `json:"jwt"`
	UserId    int64  `json:"user_id"`
	Mobile    string `json:"mobile"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
}

func (t UserLogic) Login(req *pb.LoginRequest) (res *LoginResponse, err error) {
	cli, err := t.grpcService.NewUserRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.Login(context.Background(), req)
	if err != nil {
		return nil, err
	}
	jwt, err := jwtx.GenerateJwt(resp.Id)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Jwt:       jwt,
		UserId:    resp.Id,
		Mobile:    resp.Mobile,
		Nickname:  resp.Nickname,
		AvatarUrl: resp.AvatarUrl,
	}, nil
}
