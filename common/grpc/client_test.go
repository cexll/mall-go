package grpc

import (
	"context"
	"fmt"
	bpb "mall-go/app/balance/pb"
	"mall-go/app/user/pb"
	"testing"
)

func TestNewClient(t *testing.T) {
	cli, err := NewClient(":10002", func(options *ClientOptions) {})
	if err != nil {
		fmt.Println(err)
		return
	}

	ucli := pb.NewUserClient(cli.Conn())
	uresp, err := ucli.Register(context.Background(), &pb.RegisterRequest{
		Mobile:   "19992399999",
		Nickname: "cexll",
		Password: "123123",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(uresp)

	cli, err = NewClient(":10012", func(options *ClientOptions) {})
	if err != nil {
		fmt.Println(err)
		return
	}
	bcli := bpb.NewBalanceClient(cli.Conn())
	bresp, err := bcli.GetBalance(context.Background(), &bpb.GetBalanceRequest{Type: 1, UserId: 2})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bresp)

}
