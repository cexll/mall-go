package logic

import (
	"context"
	"mall-go/app/community/cmd/api/internal/svc"
	"mall-go/app/community/cmd/pb"
)

type PushPostLogic struct {
}

func (l *PushPostLogic) PushPost(in *pb.PushPostRequest) (*pb.PushPostResponse, error) {
	resp, err := svc.Context.CommunityRpc.PushPost(context.Background(), in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
