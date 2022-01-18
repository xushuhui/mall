package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	app "mall-go/api/app/service"
	sku "mall-go/api/sku/service"
	"mall-go/app/mall/interface/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"

	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBannerRepo, NewThemeRepo, NewAppServiceClient, NewRegistrar,
	NewDiscovery, NewActivityRepo, NewCategoryRepo,
)

// Data .
type Data struct {
	ac  app.AppClient
	sc  sku.SkuClient
	log *log.Helper
	//rdb *redis.Client
}

// NewData .
func NewData(ac app.AppClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		ac:  ac,
		log: log.NewHelper(logger),
	}, cleanup, nil
}
func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
func NewAppServiceClient(r registry.Discovery) app.AppClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///app.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := app.NewAppClient(conn)
	return c
}
