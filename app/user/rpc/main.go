package main

import (
	"github.com/mix-go/dotenv"
	"github.com/mix-go/xcli"
	"rpc/commands"
	_ "rpc/config/dotenv"
	_ "rpc/config/viper"
	_ "rpc/di"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
