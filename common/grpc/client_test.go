package grpc

import (
	"context"
	"fmt"
	bpb "mall-go/app/balance/pb"
	"mall-go/app/user/pb"
	"testing"
)

func TestNewClient(t *testing.T) {
	userConn, err := NewClient(UserGRPCAddr, func(options *ClientOptions) {})
	if err != nil {
		fmt.Println(err)
		return
	}

	userCli := pb.NewUserClient(userConn.Conn())
	registerResp, err := userCli.Register(context.Background(), &pb.RegisterRequest{
		Mobile:   "19992399999",
		Nickname: "cexll",
		Password: "123123",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(registerResp)

	balanceConn, err := NewClient(BalanceGRPCAddr, func(options *ClientOptions) {})
	if err != nil {
		fmt.Println(err)
		return
	}
	balanceCli := bpb.NewBalanceClient(balanceConn.Conn())
	getBalanceResp, err := balanceCli.GetBalance(context.Background(), &bpb.GetBalanceRequest{Type: 1, UserId: 2})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getBalanceResp)

	getUserResp, err := userCli.GetUser(context.Background(), &pb.GetUserRequest{
		Id: 1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getUserResp)
}
