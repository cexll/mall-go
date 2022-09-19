package conf

type ApiConf struct {
	ServiceConf `mapstructure:",squash"`
	Addr        string `json:"Addr,default=:8080"`
	Mode        string `json:"Mode,default=release"`
	Timeout     int64  `json:"Timeout,default=60"`
}

type ServiceConf struct {
	Name       string `json:"Name"`
	Mode       string `json:"Mode,default=release"`
	MetricsUrl string `json:"MetricsUrl"`
	Prometheus string `json:"Prometheus"`
}
