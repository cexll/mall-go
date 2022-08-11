package main

import (
	"github.com/mix-go/xcli"
	"mall-go/app/goods/cmd/api/commands"
	"mall-go/common/config"
	_ "mall-go/common/di"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(config.Conf.APPDEBUG)
	xcli.AddCommand(commands.Commands...).Run()
}
