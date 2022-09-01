package logic

import (
	"context"
	balancepb "mall-go/app/balance/cmd/pb"
	"mall-go/app/user/cmd/api/services"
	"mall-go/app/user/cmd/pb"
	"mall-go/common/jwtx"
)

type UserLogic struct {
	grpcService services.GRPCService
}

func (t *UserLogic) Register(req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
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

func (t *UserLogic) GetUser(req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
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

func (t *UserLogic) SetUser(req *pb.SetUserRequest) (*pb.SetUserResponse, error) {
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

func (t *UserLogic) Logout(req *pb.LogOutRequest) (*pb.LogOutResponse, error) {
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

func (t *UserLogic) Login(req *pb.LoginRequest) (res *LoginResponse, err error) {
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

type GetBalanceResponse struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Type      int32  `json:"type"`
	Available uint64 `json:"available" default:"0"`
	Frozen    uint64 `json:"frozen" default:"0"`
}

func (t *UserLogic) GetBalance(in *balancepb.GetBalanceRequest) (*GetBalanceResponse, error) {
	cli, err := t.grpcService.NewBalanceRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.GetBalance(context.Background(), in)
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

func (t *UserLogic) GetBalanceChangeList(in *balancepb.GetBalanceChangeListRequest) (*balancepb.GetBalanceChangeListResponse, error) {
	cli, err := t.grpcService.NewBalanceRpcClient()
	if err != nil {
		return nil, err
	}
	resp, err := cli.GetBalanceChangeList(context.Background(), in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
