package data

import (
	"context"
	"mall-go/app/mall/interface/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"

	"mall-go/api/mall"
	appv1 "mall-go/api/mall"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBannerRepo, NewAppServiceClient,NewRegistrar,
	NewDiscovery,
)

// Data .
type Data struct {
	ac  mall.AppClient
	log *log.Helper
	//rdb *redis.Client
}

// NewData .
func NewData(ac mall.AppClient, logger log.Logger) (*Data, func(), error) {
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
func NewAppServiceClient(r registry.Discovery) appv1.AppClient {
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
	c := appv1.NewAppClient(conn)
	return c
}
