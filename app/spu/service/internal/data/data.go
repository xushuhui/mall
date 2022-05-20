package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"mall-go/app/spu/service/internal/conf"
	"time"
)

var ProviderSet = wire.NewSet(NewData, NewMongo, NewSkuRepo, NewSpuRepo)

// Data .
type Data struct {
	log *log.Helper
	db  *mongo.Database
}

// NewData .
func NewData(db *mongo.Database, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  db,
		log: log.NewHelper(logger),
	}, cleanup, nil
}

func NewMongo(conf *conf.Data) *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongodb.Uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client.Database(conf.Mongodb.Database)
}
