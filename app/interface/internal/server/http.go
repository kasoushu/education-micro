package server

import (
	iv1 "education/api/v1/interface"
	"education/app/interface/internal/conf"
	"education/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHttpServer(c *conf.AppConfig, userSvc *service.InterfaceService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			//jwt.Server(func(token *jwt2.Token) (interface{}, error) {
			//	return c.Auth.ServiceKey, nil
			//}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	iv1.RegisterEducationInterfaceHTTPServer(srv, userSvc)
	return srv
}
