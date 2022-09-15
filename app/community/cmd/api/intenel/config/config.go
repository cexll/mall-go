package config

import "mall-go/common/config"

type Config struct {
	JwtAuth struct {
		Secret string
	}
	Api config.ApiConf
}
