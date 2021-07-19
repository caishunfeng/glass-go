package logic

import (
	"context"
	"errors"
	"fmt"
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

	newHost, err := redisClient.Get(getProxyKey(host))
	if err != nil || newHost == "" {
		return err
	}

	if newHost == "" {
		return errors.New("cat not get proxy domain")
	}

	logx.Infof("proxy host:%s, newHost:%s", host, newHost)

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

func (l *ProxyLogic) Registry(req types.DomainRegistryReq) error {
	redisCfg := config.GetConfig().Redis
	redisClient := redis.NewRedis(redisCfg.Addr, redisCfg.Type, redisCfg.Pass)

	if req.From == "" {
		return errors.New("from is empty")
	}

	if req.To == "" {
		return errors.New("to is empty")
	}

	if req.Expire == 0 {
		req.Expire = config.GetConfig().Domain.DefaultExpire
	}

	var val bool
	var err error
	if req.Expire == NoExpire {
		val, err = redisClient.Setnx(getProxyKey(req.From), req.To)
	} else {
		val, err = redisClient.SetnxEx(getProxyKey(req.From), req.To, config.GetConfig().Domain.DefaultExpire)
	}
	if err != nil {
		return err
	}
	if val != true {
		return errors.New("domain had exist")
	}
	return nil
}

func (l *ProxyLogic) InternalRegistry(r *http.Request, req types.InternalDomainRegistryReq) error {
	redisCfg := config.GetConfig().Redis
	redisClient := redis.NewRedis(redisCfg.Addr, redisCfg.Type, redisCfg.Pass)

	ip := clientIP(r)
	if ip == "" {
		return errors.New("get ip error")
	}

	if req.Domain == "" {
		return errors.New("domain is empty")
	}

	if req.LocalPort == 0 {
		return errors.New("local port is empty")
	}

	if req.Expire == 0 {
		req.Expire = config.GetConfig().Domain.DefaultExpire
	}

	var val bool
	var err error
	if req.Expire == NoExpire {
		val, err = redisClient.Setnx(getProxyKey(req.Domain), fmt.Sprintf("%s:%d", ip, req.LocalPort))
	} else {
		val, err = redisClient.SetnxEx(getProxyKey(req.Domain), fmt.Sprintf("%s:%d", ip, req.LocalPort), config.GetConfig().Domain.DefaultExpire)
	}

	if err != nil {
		return err
	}
	if val != true {
		return errors.New("domain had exist")
	}
	return nil
}

func getProxyKey(key string) string {
	return config.GetConfig().Domain.Prefix + key
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
