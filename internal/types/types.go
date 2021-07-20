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
	From   string `form:"from"`            // 源域名地址
	To     string `form:"to"`              // 目标域名地址
	Expire int    `form:"expire,optional"` // 过期时间, 为0或者空则使用配置的默认时间，-1代表永久
}

type DomainCancelReq struct {
	From string `form:"from"`
}
