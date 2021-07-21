package middleware

import (
	"net/http"
)

type RespMiddleware struct {
}

func NewRespMiddleware() *RespMiddleware {
	return &RespMiddleware{}
}

func (m *RespMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")             //返回数据格式是json
		next(w, r)
	}
}
