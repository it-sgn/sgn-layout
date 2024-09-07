package server

import (
	"context"

	exampleV1 "github.com/it-sgn/sgn-layout/api/example/v1"
	"github.com/it-sgn/sgn-layout/internal/conf"
	"github.com/it-sgn/sgn-layout/internal/service"
	"github.com/it-sgn/sgn-layout/pkg/middleware/requestInfo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// func initTracerProvider(ctx context.Context, res *resource.Resource, conn string) (func(context.Context) error, error) {
// 	// Set up a trace exporter

// 	// 服务端的Jaeger支持HTTPS时使用
// 	// traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(conn))

// 	// 服务端的Jaeger不支持HTTPS时使用otlptracehttp.WithInsecure()显式声明只使用HTTP不安全的连接
// 	traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithInsecure(), otlptracehttp.WithEndpoint(conn))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
// 	}

// 	// Register the trace exporter with a TracerProvider, using a batch
// 	// span processor to aggregate spans before export.
// 	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
// 	tracerProvider := sdktrace.NewTracerProvider(
// 		// sdktrace.WithSampler(sdktrace.AlwaysSample()),
// 		// 将基于父span的采样率设置
// 		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1.0))),
// 		// 始终确保在生产中批量处理
// 		sdktrace.WithBatcher(traceExporter),
// 		// 在资源中记录有关此应用程序的信息
// 		sdktrace.WithResource(res),
// 		sdktrace.WithSpanProcessor(bsp),
// 	)
// 	otel.SetTracerProvider(tracerProvider)

// 	// Set global propagator to tracecontext (the default is no-op).
// 	otel.SetTextMapPropagator(propagation.TraceContext{})

// 	// Shutdown will flush any remaining spans and shut down the exporter.
// 	return tracerProvider.Shutdown, nil
// }

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	conf *conf.Bootstrap,
	server *conf.Server,
	tr *conf.Trace,
	service *service.ExampleService,
	logger log.Logger) *http.Server {

	ctx := context.Background()
	// fmt.Println(bc.con)
	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			// serviceName,
			semconv.ServiceNameKey.String("kratos-trace"),
			attribute.String("exporter", "example-http"),
			attribute.String("environment", "dev"),
			attribute.Float64("float", 312.23),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	// shutdownTracerProvider, err := initTracerProvider(ctx, res, tr.Jaeger.Http.Endpoint)
	_, err2 := initTracerProvider(ctx, res, tr.Jaeger.Endpoint)
	if err2 != nil {
		log.Fatal(err)
	}

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			metadata.Server(),
			tracing.Server(), // trace
			// Access log
			logging.Server(logger),

			// Tenants
			requestInfo.SetRequestInfo(),
		),
	}
	if server.Http.Network != "" {
		opts = append(opts, http.Network(server.Http.Network))
	}
	if server.Http.Addr != "" {
		opts = append(opts, http.Address(server.Http.Addr))
	}
	if server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	exampleV1.RegisterExampleServiceHTTPServer(srv, service)
	return srv
}
