package main

import (
	commands "github.com/cexll/mall-go/commands"
	_ "github.com/cexll/mall-go/config/dotenv"
	_ "github.com/cexll/mall-go/config/viper"
	_ "github.com/cexll/mall-go/di"
	"github.com/mix-go/dotenv"
	"github.com/mix-go/xcli"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
