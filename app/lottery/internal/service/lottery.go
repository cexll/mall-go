package service

import (
	"context"
	lottery "mall-go/app/lottery/api/lottery/v1"
)

func (l *LotteryService) FindLottery(ctx context.Context, req *lottery.FindLotteryReq) (*lottery.LotteryResp, error) {
	return &lottery.LotteryResp{
		Id:          req.Id,
		Name:        "x",
		Description: "x",
	}, nil
}

func (l *LotteryService) CreateLottery(ctx context.Context, req *lottery.CreateLotteryReq) (*lottery.LotteryResp, error) {
	return &lottery.LotteryResp{
		Id:          1,
		Name:        req.Name,
		Description: req.Description,
	}, nil
}
