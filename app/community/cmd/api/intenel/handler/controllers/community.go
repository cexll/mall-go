package controllers

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"mall-go/app/community/cmd/api/intenel/svc"
	"mall-go/app/community/cmd/pb"
)

type CommunityController struct {
}

func (t *CommunityController) PushPost(ctx context.Context, c *app.RequestContext) {
	in := pb.PushPostRequest{}
	resp, err := svc.Context.CommunityRpc.PushPost(ctx, &in)
	if err != nil {
		//TODO
	}
	fmt.Println(resp)
	c.String(200, "hello")
}
