//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"mall-go/app/sku/service/internal/biz"
	"mall-go/app/sku/service/internal/conf"
	"mall-go/app/sku/service/internal/data"
	"mall-go/app/sku/service/internal/server"
	"mall-go/app/sku/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
