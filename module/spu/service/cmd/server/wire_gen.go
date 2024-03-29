// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"mall-go/module/spu/service/internal/biz"
	"mall-go/module/spu/service/internal/conf"
	"mall-go/module/spu/service/internal/data"
	"mall-go/module/spu/service/internal/server"
	"mall-go/module/spu/service/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	database := data.NewMongo(confData)
	dataData, cleanup, err := data.NewData(database, logger)
	if err != nil {
		return nil, nil, err
	}
	skuRepo := data.NewSkuRepo(dataData, logger)
	skuUsecase := biz.NewSkuUsecase(skuRepo, logger)
	spuRepo := data.NewSpuRepo(dataData, logger)
	spuUsecase := biz.NewSpuUsecase(spuRepo, logger)
	spuService := service.NewSpuService(skuUsecase, spuUsecase)
	httpServer := server.NewHTTPServer(confServer, spuService, logger)
	grpcServer := server.NewGRPCServer(confServer, spuService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
