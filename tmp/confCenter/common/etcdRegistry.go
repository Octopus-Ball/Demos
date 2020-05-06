package common

import (
	"time"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

// GetEtcdRegistry 获取EtcdRegistry
func GetEtcdRegistry() registry.Registry {
	etcdDriver := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = GetEtcdConfig().Endpoints
		op.Timeout = time.Duration(GetEtcdConfig().DialTimeout) * time.Second
	})
	return etcdDriver
}