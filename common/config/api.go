package config

type ApiConf struct {
	Addr string `json:"addr,default=:8080"`
	Mode string `json:"mode,default=release"`
}
