package server

import (
	cv1 "education/api/v1/course"
	"education/app/selectCource/internal/conf"
	"education/app/selectCource/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.AppConfig, service *service.CourseService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(log.With(logger,
				//"ts", log.DefaultTimestamp,
				"caller", log.DefaultCaller,
				//"interface.id", id,
				"trace_id", tracing.TraceID(),
				"span_id", tracing.SpanID(),
			)),
		),
	}
	//log.NewHelper(logger).Debug("grpc init!")
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
	cv1.RegisterCourseServer(srv, service)
	return srv
}
