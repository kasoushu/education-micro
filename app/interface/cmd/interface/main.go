package main

import (
	"education/app/interface/internal/conf"
	"education/pkg"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// go build -ldflags "-X main.Version=x.y.z"

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "education.interface"
	// Version is the version of the compiled software.
	Version string = "0.1.0"
	// flagconf is the config flag.
	flagconf string
	id       = uuid.New().String()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}
func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rs registry.Registrar) *kratos.App {
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
		kratos.Registrar(rs),
	)
}
func loggerInit() log.Logger {
	cusLog := pkg.NewLogger()
	logger := log.With(cusLog) //"ts", log.DefaultTimestamp,
	//"caller", log.DefaultCaller,
	//"interface.id", id,
	//"trace_id", tracing.TraceID(),
	//"span_id", tracing.SpanID(),
	log.NewHelper(logger).Info("interface is initiating!")
	log.NewHelper(logger).Infof("Service Name:\x1b[31m%s\x1b[0m  \x1b[34mService Version:\x1b[0m \x1B[32m%s\x1B[0m", Name, Version)
	log.SetLogger(logger)
	return logger
}
func main() {

	logger := loggerInit()
	// tracing provider
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}
	var appConfig conf.AppConfig
	if err := c.Scan(&appConfig); err != nil {
		panic(err)
	}
	// init tracing
	SetTracerProvider(appConfig.Jaeger.Address, logger)
	app, cleanup, err := initApp(&appConfig, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
func SetTracerProvider(c string, logger log.Logger) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c)))
	if err != nil {
		panic(err)
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in an Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(Name),
			//attribute.String("environment", en),
			attribute.String("UUID", id),
		)),
	)
	log.NewHelper(logger).Info("jaeger init successful!")
	otel.SetTracerProvider(tp)
}
