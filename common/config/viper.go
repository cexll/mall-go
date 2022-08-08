package config

import (
	"github.com/mix-go/xcli/flag"
	"github.com/spf13/viper"
	"mall-go/common"
)

type Configure struct {
	APPDEBUG bool
	API      APIConfig
	RPC      RPConfig
	JWT      JWTConfig
	REDIS    REDISConfig
	DB       DBConfig
}

var Conf Configure

type APIConfig struct {
	GinAddr string
	GinMode string
}

type RPConfig struct {
	GRPCAddr string
}

type JWTConfig struct {
	// HMAC_SECRET
	HmacSecret string
}

type REDISConfig struct {
	RedisAddr     string
	RedisPassword string
	RedisDatabase int64
	RedisTimeOut  int64
}

type DBConfig struct {
	DSN string
}

func init() {
	// Conf support JSON, TOML, YAML, HCL, INI, envfile
	path := flag.Match("f", "config").String("etc/config.yaml")

	v := viper.New()

	viper.SetConfigType("yaml")
	path = common.GetCWD() + "/" + path

	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic(err)
	}
}
