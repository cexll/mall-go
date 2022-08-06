package main

import (
	"github.com/mix-go/dotenv"
	"github.com/mix-go/xcli"
	"mall-go/app/balance/cmd/api/commands"
	_ "mall-go/app/balance/cmd/api/config/dotenv"
	_ "mall-go/app/balance/cmd/api/config/viper"
	_ "mall-go/common/di"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}