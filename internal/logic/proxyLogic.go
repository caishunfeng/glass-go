package logic

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"glass/internal/config"
	"glass/internal/svc"
	"glass/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type ProxyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const NoExpire int = -1

func NewProxyLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProxyLogic {
	return ProxyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProxyLogic) Proxy(w http.ResponseWriter, r *http.Request, req types.ProxyReq) error {
	r.ParseForm()

	host := r.Host

	r.RequestURI = req.Path
	r.URL.Path = req.Path

	redisCfg := config.GetConfig().Redis
	redisClient := redis.NewRedis(redisCfg.Addr, redisCfg.Type, redisCfg.Pass)

	newHost, err := redisClient.Hget(config.GetConfig().Domain.Prefix, host)
	if err != nil || newHost == "" {
		return err
	}

	if newHost == "" {
		return errors.New("cat not get proxy domain")
	}

	logx.Infof("proxy log -> host:%s, newHost:%s", host, newHost)

	remoteUrl := "http://" + newHost

	remote, err := url.Parse(remoteUrl)
	if err != nil {
		return err
	}

	// 转发
	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ServeHTTP(w, r)

	return nil
}

func (l *ProxyLogic) Registry(req types.DomainRegistryReq) (bool, error) {
	redisCfg := config.GetConfig().Redis
	redisClient := redis.NewRedis(redisCfg.Addr, redisCfg.Type, redisCfg.Pass)

	if req.From == "" {
		return false, errors.New("from is empty")
	}

	if req.To == "" {
		return false, errors.New("to is empty")
	}

	if req.Expire == 0 {
		req.Expire = config.GetConfig().Domain.DefaultExpire
	}

	val, err := redisClient.Hsetnx(config.GetConfig().Domain.Prefix, req.From, req.To)
	if err != nil {
		return false, err
	}
	if val != true {
		return false, errors.New("domain had exist")
	}

	if req.Expire != NoExpire {
		return val, redisClient.Expire(config.GetConfig().Domain.Prefix+":"+req.From, req.Expire)
	}

	return val, err
}

func (l *ProxyLogic) Cancel(req types.DomainCancelReq) (val bool, err error) {
	redisCfg := config.GetConfig().Redis
	redisClient := redis.NewRedis(redisCfg.Addr, redisCfg.Type, redisCfg.Pass)
	return redisClient.Hdel(config.GetConfig().Domain.Prefix, req.From)
}

func (l *ProxyLogic) GetProxyKeys() (val []types.ProxyKeyResp, err error) {
	redisCfg := config.GetConfig().Redis
	redisClient := redis.NewRedis(redisCfg.Addr, redisCfg.Type, redisCfg.Pass)
	res, err := redisClient.Hgetall(config.GetConfig().Domain.Prefix)
	if err != nil {
		return nil, err
	}
	for k, v := range res {
		val = append(val, types.ProxyKeyResp{From: k, To: v})
	}
	return
}

func clientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
