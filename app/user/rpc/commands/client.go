package commands

import (
	"context"
	"fmt"
	"github.com/mix-go/dotenv"
	"google.golang.org/grpc"
	"mall-go/app/user/pb"
	"time"
)

type GrpcClientCommand struct {
}

func (t *GrpcClientCommand) Main() {
	addr := dotenv.Getenv("GIN_ADDR").String(":8080")
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewUserClient(conn)
	req := pb.GetUserRequest{
		Id: 80,
	}
	resp, err := cli.GetUser(ctx, &req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("User: %s", resp.String()))
}
