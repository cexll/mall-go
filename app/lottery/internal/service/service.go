package service

import (
	"mall-go/app/lottery/internal/svc"

	lottery "mall-go/app/lottery/api/lottery/v1"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewLotteryService)

type LotteryService struct {
	lottery.UnimplementedLotteryServer

	svcCtx *svc.ServiceContext
}

func NewLotteryService(ctx *svc.ServiceContext) *LotteryService {
	return &LotteryService{
		svcCtx: ctx,
	}
}
