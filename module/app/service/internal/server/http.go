package server

import (
	"mall-go/module/app/service/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/validate"

	"mall-go/module/app/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, as *service.AppService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			validate.Validator(),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	service.RegisterAppHTTPServer(srv, as)

	return srv
}
