package conf

type RedisConf struct {
	Addr     string `json:"addr,default=127.0.0.1:6379"`
	Pass     string `json:"pass,default="`
	DataBase int64  `json:"database,default=1"`
	Timeout  int64  `json:"timeout,default=60"`
}
