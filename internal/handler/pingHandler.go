package handler

import (
	"net/http"

	"glass/internal/logic"
	"glass/internal/svc"
	"glass/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func PingHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPingLogic(r.Context(), ctx)
		resp, err := l.Ping(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
