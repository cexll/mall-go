package config

type RpcClientConf struct {
	Addr    string `json:"addr"`
	Timeout int64  `json:"timeout"`
}
