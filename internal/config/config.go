package config

import (
	"fmt"
	"glass/internal/utils"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/jsonx"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
)

type RedisConfig struct {
	Addr string
	Pass string
	Type string
}

type DomainConfig struct {
	Prefix        string
	DefaultExpire int
}

type Config struct {
	Api    rest.RestConf
	Logx   logx.LogConf
	Redis  RedisConfig
	Domain DomainConfig
}

var configFile = "etc/glass-api.yaml"
var config Config

func InitConfig() {
	conf.MustLoad(configFile, &config)

	jsonb, err := jsonx.Marshal(config)
	if err != nil {
		fmt.Println("config error", err)
	}
	fmt.Println("config: ", string(jsonb))
}

func GetConfig() Config {
	return config
}

func InitLogxConfig(logxConfig logx.LogConf) {
	logxConfig.Path = utils.GetCurrentPath() + "/" + logxConfig.Path
	logx.MustSetup(logxConfig)
}
