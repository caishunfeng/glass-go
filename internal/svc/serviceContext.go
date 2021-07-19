package svc

import (
	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config rest.RestConf
}

func NewServiceContext(c rest.RestConf) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
