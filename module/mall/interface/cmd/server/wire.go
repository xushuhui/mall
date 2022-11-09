//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mall-go/module/mall/interface/internal/biz"
	"mall-go/module/mall/interface/internal/conf"
	"mall-go/module/mall/interface/internal/data"
	"mall-go/module/mall/interface/internal/server"
	"mall-go/module/mall/interface/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, *conf.App, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
