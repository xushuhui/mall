package data

import (
	"mall-go/internal/conf"
	"mall-go/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewBannerRepo, NewActivityRepo, NewCouponRepo,
	NewThemeRepo, NewSpuRepo, NewUserRepo, NewOrderRepo)

// Data .
type Data struct {
	db  *model.Client
	log *log.Helper
	//rdb *redis.Client
}

func NewEntClient(conf *conf.Data, logger log.Logger) *model.Client {
	l := log.NewHelper(logger)
	client, err := model.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	client.Debug()
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}

	return client
}

// func NewRedisClient(conf *conf.Data, logger log.Logger)*redis.Client {
// 	l := log.NewHelper(logger)
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     conf.Redis.Addr,
// 		Password: "",
// 		DB:       0,
// 	})
// 	err := rdb.Ping(context.Background()).Err()
// 	if err != nil {
// 		l.Fatalf("failed opening connection to rdb: %v", err)
// 	}
// 	return rdb
// }

// NewData .
func NewData(entClient *model.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  entClient,
		log: log.NewHelper(logger),
	}, cleanup, nil
}
