package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/encoding/protojson"
	"mall-go/pkg/utils"

	"mall-go/app/mall/interface/internal/conf"

	zaplog "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "mall.interface"
	// Version is the version of the compiled software.
	Version = "1.0"
	// flagconf is the config flag.
	flagconf string

	id = Name
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	json.MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(rr),
	)
}
func newLogger() log.Logger {
	var tops = []utils.TeeOption{
		{
			Filename: "./access.log",
			Ropt: utils.RotateOptions{
				MaxSize:    1,
				MaxAge:     10,
				MaxBackups: 10,
			},
			Lef: func(level zapcore.Level) bool {
				return level >= zapcore.InfoLevel
			},
		},
		{
			Filename: "./error.log",
			Ropt: utils.RotateOptions{
				MaxSize:    1,
				MaxAge:     10,
				MaxBackups: 10,
			},

			Lef: func(level zapcore.Level) bool {
				return level >= zapcore.WarnLevel
			},
		},
	}

	zlogger := zaplog.NewLogger(utils.NewTeeWithRotate(tops, zap.AddCaller()))

	defer func() { _ = zlogger.Sync() }()

	logger := log.With(zlogger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	return logger
}
func main() {
	flag.Parse()

	logger := newLogger()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}
	app, cleanup, err := initApp(&bc, &rc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
