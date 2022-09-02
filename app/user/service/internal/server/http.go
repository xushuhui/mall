package server

import (
	"mall-go/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"

	"mall-go/app/user/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, as *service.UserService, logger log.Logger) *http.Server {
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

	service.RegisterUserHTTPServer(srv, as)

	return srv
}
