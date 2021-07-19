// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"glass/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: PingHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/proxy/:path",
				Handler: ProxyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/registry",
				Handler: RegistryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/registry/internal",
				Handler: InternalRegistryHandler(serverCtx),
			},
		},
	)
}
