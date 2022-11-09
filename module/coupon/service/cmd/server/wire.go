//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mall-go/module/coupon/service/internal/biz"
	"mall-go/module/coupon/service/internal/conf"
	"mall-go/module/coupon/service/internal/data"
	"mall-go/module/coupon/service/internal/server"
	"mall-go/module/coupon/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
