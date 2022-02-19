package pkg

import (
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.4.0"
)

type JaegerConfig struct {
	Name string
	Url  string
	ID   string
}

func SetTracerProvider(c JaegerConfig, logger log.Logger) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(c.Url)))
	if err != nil {
		panic(err)
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in an Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(c.Name),
			//attribute.String("environment", en),
			attribute.String("UUID", c.ID),
		)),
	)
	log.NewHelper(logger).Info("jaeger init successful!")
	otel.SetTracerProvider(tp)
}
