package handler

import (
	"net/http"

	"glass/internal/logic"
	"glass/internal/svc"
	"glass/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ProxyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProxyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProxyLogic(r.Context(), ctx)
		err := l.Proxy(w, r, req)
		if err != nil {
			httpx.Error(w, err)
		}
	}
}

func RegistryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DomainRegistryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProxyLogic(r.Context(), ctx)
		err := l.Registry(req)
		if err != nil {
			httpx.Error(w, err)
		}
	}
}

func InternalRegistryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InternalDomainRegistryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProxyLogic(r.Context(), ctx)
		err := l.InternalRegistry(r, req)
		if err != nil {
			httpx.Error(w, err)
		}
	}
}
