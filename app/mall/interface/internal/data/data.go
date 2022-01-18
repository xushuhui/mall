package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	app "mall-go/api/app/service"
	sku "mall-go/api/spu/service"
	user "mall-go/api/user/service"
	"mall-go/app/mall/interface/internal/conf"

	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBannerRepo, NewThemeRepo, NewRegistrar,
	NewDiscovery, NewActivityRepo, NewCategoryRepo, NewAppServiceClient, NewSkuServiceClient, NewUserRepo, NewWeappClient, NewUserServiceClient,
)

// Data .
type Data struct {
	ac  app.AppClient
	sc  sku.SkuClient
	uc  user.UserClient
	wc  *WeappClient
	log *log.Helper
	//rdb *redis.Client
}
type WeappClient struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

func NewWeappClient(c *conf.Bootstrap) *WeappClient {
	return &WeappClient{
		Appid:  c.Weapp.Appid,
		Secret: c.Weapp.Sercret,
	}
}

// NewData .
func NewData(ac app.AppClient, sc sku.SkuClient, uc user.UserClient, wc *WeappClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		ac:  ac,
		sc:  sc,
		wc:  wc,
		uc:  uc,
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
func NewUserServiceClient(r registry.Discovery) user.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := user.NewUserClient(conn)
	return c
}
func NewSkuServiceClient(r registry.Discovery) sku.SkuClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///sku.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := sku.NewSkuClient(conn)
	return c
}
