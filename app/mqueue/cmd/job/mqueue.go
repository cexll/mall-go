package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"mall-go/app/mqueue/cmd/job/internal/config"
	"mall-go/app/mqueue/cmd/job/internal/logic"
	"mall-go/app/mqueue/cmd/job/internal/svc"
	"mall-go/common/conf"
	"os"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
