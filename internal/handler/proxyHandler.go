package handler

import (
	"net/http"

	"glass/internal/base"
	"glass/internal/logic"
	"glass/internal/svc"
	"glass/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ProxyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProxyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
			return
		}

		l := logic.NewProxyLogic(r.Context(), ctx)
		err := l.Proxy(w, r, req)
		if err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
		}
	}
}

func RegistryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DomainRegistryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
			return
		}

		l := logic.NewProxyLogic(r.Context(), ctx)
		val, err := l.Registry(req)
		if err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
		}
		httpx.OkJson(w, base.NewResult(base.OK, val))
	}
}

func GetProxyKeysHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewProxyLogic(r.Context(), ctx)
		val, err := l.GetProxyKeys()
		if err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
		}
		httpx.OkJson(w, base.NewResult(base.OK, val))
	}
}

func CancelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DomainCancelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
			return
		}

		l := logic.NewProxyLogic(r.Context(), ctx)
		val, err := l.Cancel(req)
		if err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
		}
		httpx.OkJson(w, base.NewResult(base.OK, val))
	}
}
