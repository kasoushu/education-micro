package server

import (
	iv1 "education/api/v1/interface"
	"education/app/interface/internal/conf"
	"education/app/interface/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.AppConfig, userSvc *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return c.Auth.ServiceKey, nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	iv1.RegisterInterfaceServer(srv, userSvc)
	return srv
}
