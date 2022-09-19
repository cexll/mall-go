package controllers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"mall-go/app/community/cmd/api/internal/svc"
	"mall-go/app/community/cmd/pb"
)

type CommunityController struct {
}

func (t *CommunityController) PushPost(ctx context.Context, c *app.RequestContext) {
	in := pb.PushPostRequest{}
	resp, err := svc.Context.CommunityRpc.PushPost(ctx, &in)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}
