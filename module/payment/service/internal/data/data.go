package data

import (
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewPayRepo)

// Data .
type Data struct {
	log *log.Helper
	// rdb *redis.Client
}

// NewData .
func NewData(logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		log: log.NewHelper(logger),
	}, cleanup, nil
}
