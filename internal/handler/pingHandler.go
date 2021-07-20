package handler

import (
	"net/http"

	"glass/internal/base"
	"glass/internal/logic"
	"glass/internal/svc"
	"glass/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func PingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
			return
		}

		l := logic.NewPingLogic(r.Context(), ctx)
		resp, err := l.Ping(req)
		if err != nil {
			httpx.Error(w, base.NewError(base.FAIL, err.Error()))
		} else {
			httpx.OkJson(w, base.NewResult(base.OK, resp))
		}
	}
}
