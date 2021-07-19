// Code generated by goctl. DO NOT EDIT.
package types

type PingReq struct {
}

type PongResp struct {
	T int64 `json:"t"`
}

type ProxyReq struct {
	Path string `path:"path"`
}

type ProxyResp struct {
}

type DomainRegistryReq struct {
	From   string `form:"from"`   // 源域名地址
	To     string `form:"to"`     // 目标域名地址
	Expire int    `form:"expire"` // 过期时间, 为0或者空则使用配置的默认时间，-1代表永久
}

type InternalDomainRegistryReq struct {
	Domain    string `form:"domain"` // 注册域名
	LocalPort int    `form:"port"`   // 本地端口
	Expire    int    `form:"expire"` // 过期时间, 为0或者空则使用配置的默认时间，-1代表永久
}
