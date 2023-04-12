package main

import (
	"flag"

	"mall-go/app/lottery/internal/config"

	"github.com/mix-plus/core/conf"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	app, err := initApp(&c)
	if err != nil {
		panic(err)
	}

	app.Run()
}
