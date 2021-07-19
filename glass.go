package main

import (
	"flag"

	"glass/internal/config"
	"glass/internal/handler"
	"glass/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/proc"
	"github.com/tal-tech/go-zero/rest"
)

func main() {
	flag.Parse()

	config.InitConfig()

	// 日志初始化
	config.InitLogxConfig(config.GetConfig().Logx)

	// 服务初始化
	apiConfig := config.GetConfig().Api
	ctx := svc.NewServiceContext(apiConfig)
	server := rest.MustNewServer(apiConfig)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	logx.Infof("Starting server at %s:%d...\n", apiConfig.Host, apiConfig.Port)

	server.Start()

	proc.AddShutdownListener(func() {
		logx.Close()
	})

}
