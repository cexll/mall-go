package main

import (
	"github.com/mix-go/dotenv"
	"github.com/mix-go/xcli"
	"mall-go/app/user/cmd/rpc/commands"
	_ "mall-go/app/user/cmd/rpc/config/dotenv"
	_ "mall-go/app/user/cmd/rpc/config/viper"
	_ "mall-go/common/di"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}