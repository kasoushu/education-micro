package main

import (
	"education/app/selectCource/internal/conf"
	"education/pkg"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/uuid"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "education.select.course"
	// Version is the version of the compiled software.
	Version string = "0.1.0"
	// flagconf is the config flag.
	flagconf string

	id = uuid.New().String()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, conf *conf.AppConfig, gs *grpc.Server, rs registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rs),
	)
}
func loggerInit() log.Logger {
	cusLog := pkg.NewLogger()
	logger := log.With(cusLog)//"ts", log.DefaultTimestamp,
	//"caller", log.DefaultCaller,
	//"interface.id", id,

	//"trace_id", tracing.TraceID(),
	//"span_id", tracing.SpanID(),

	log.NewHelper(logger).Info("service is initiating!")
	log.NewHelper(logger).Infof("Service Name:\x1b[31m%s\x1b[0m  \x1b[34mService Version:\x1b[0m \x1B[32m%s\x1B[0m", Name, Version)
	log.SetLogger(logger)
	return logger
}
func main() {
	//logger init
	logger := loggerInit()
	// conifg load
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithLogger(logger),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}
	var myConfig conf.AppConfig
	if err := c.Scan(&myConfig); err != nil {
		panic(err)
	}

	//tracing
	pkg.SetTracerProvider(pkg.JaegerConfig{
		Name: Name,
		Url:  myConfig.Jaeger.Address,
		ID:   id,
	}, logger)

	//app, cleanup, err := initApp(myConfig.Server, myConfig.Data, logger)
	app, cleanup, err := initApp(&myConfig, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
