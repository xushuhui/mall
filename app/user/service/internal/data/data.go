package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"mall-go/app/user/service/internal/conf"
	"mall-go/app/user/service/internal/data/model"
)

var ProviderSet = wire.NewSet(NewData, NewEntClient, NewUserRepo)

// Data .
type Data struct {
	db  *model.Client
	log *log.Helper
	// rdb *redis.Client
}

func NewEntClient(conf *conf.Data, logger log.Logger) *model.Client {
	l := log.NewHelper(logger)
	client, err := model.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}
	client = client.Debug()

	// if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
	// 	l.Fatalf("failed creating schema resources: %v", err)
	// }
	return client
}

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

func WithTx(ctx context.Context, client *model.Client, fn func(tx *model.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
