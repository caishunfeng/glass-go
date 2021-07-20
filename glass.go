package main

import (
	"flag"
	"net/http"

	"glass/internal/base"
	"glass/internal/config"
	"glass/internal/handler"
	"glass/internal/middleware"
	"glass/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/proc"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
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

	// 中间件
	server.Use(middleware.NewRespMiddleware().Handle)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *base.CodeError:
			return http.StatusOK, e.Error()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	logx.Infof("Starting server at %s:%d...\n", apiConfig.Host, apiConfig.Port)

	server.Start()

	proc.AddShutdownListener(func() {
		logx.Close()
	})

}
