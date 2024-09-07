package server

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

// func initTracer(ctx context.Context, url string) error {
// 	// 创建 Jaeger exporter
// 	// exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
// 	otlpExp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(url))
// 	if err != nil {
// 		return err
// 	}
// 	tp := tracesdk.NewTracerProvider(
// 		// 将基于父span的采样率设置为100%
// 		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
// 		// 始终确保在生产中批量处理
// 		tracesdk.WithBatcher(otlpExp),
// 		// tracesdk.WithBatcher(exp),
// 		// 在资源中记录有关此应用程序的信息
// 		tracesdk.WithResource(resource.NewSchemaless(
// 			semconv.ServiceNameKey.String("kratos-trace"),
// 			attribute.String("exporter", "jaeger"),
// 			attribute.Float64("float", 312.23),
// 		)),
// 	)
// 	otel.SetTracerProvider(tp)
// 	return nil
// }

func initTracerProvider(ctx context.Context, res *resource.Resource, conn string) (func(context.Context) error, error) {
	// Set up a trace exporter

	// Used when Jaeger on the server side supports HTTPS
	// traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(conn))

	// When the Jaeger on the server does not support HTTPS, use otlptracehttp.WithInsecure() to explicitly declare that only HTTP insecure connections will be used.
	traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithInsecure(), otlptracehttp.WithEndpoint(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		// sdktrace.WithSampler(sdktrace.AlwaysSample()),
		// Set the sample rate based on the parent span
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1.0))),
		// Always ensure batch processing in production
		sdktrace.WithBatcher(traceExporter),
		// Document information about this application in resources
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// Set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tracerProvider.Shutdown, nil
}

// func initTracerHttp(ctx context.Context, res *resource.Resource, conn string) (func(context.Context) error, error) {
// 	// Set up a trace exporter

// 	// Used when Jaeger on the server side supports HTTPS
// 	// traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(conn))

// 	// When the Jaeger on the server does not support HTTPS, use otlptracehttp.WithInsecure() to explicitly declare that only HTTP insecure connections will be used.
// 	traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithInsecure(), otlptracehttp.WithEndpoint(conn))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
// 	}

// 	// Register the trace exporter with a TracerProvider, using a batch
// 	// span processor to aggregate spans before export.
// 	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
// 	tracerProvider := sdktrace.NewTracerProvider(
// 		// sdktrace.WithSampler(sdktrace.AlwaysSample()),
// 		// Set the sample rate based on the parent span
// 		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1.0))),
// 		// Always ensure batch processing in production
// 		sdktrace.WithBatcher(traceExporter),
// 		// Document information about this application in resources
// 		sdktrace.WithResource(res),
// 		sdktrace.WithSpanProcessor(bsp),
// 	)
// 	otel.SetTracerProvider(tracerProvider)

// 	// Set global propagator to tracecontext (the default is no-op).
// 	otel.SetTextMapPropagator(propagation.TraceContext{})

// 	// Shutdown will flush any remaining spans and shut down the exporter.
// 	return tracerProvider.Shutdown, nil
// }
// func initTracerGRPC(ctx context.Context, res *resource.Resource, conn string) (func(context.Context) error, error) {
// 	// Set up a trace exporter

// 	// Used when Jaeger on the server side supports HTTPS
// 	// traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(conn))

// 	// When the Jaeger on the server does not support HTTPS, use otlptracehttp.WithInsecure() to explicitly declare that only HTTP insecure connections will be used.
// 	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(conn))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
// 	}

// 	// Register the trace exporter with a TracerProvider, using a batch
// 	// span processor to aggregate spans before export.
// 	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
// 	tracerProvider := sdktrace.NewTracerProvider(
// 		// sdktrace.WithSampler(sdktrace.AlwaysSample()),
// 		// Set the sample rate based on the parent span
// 		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1.0))),
// 		// Always ensure batch processing in production
// 		sdktrace.WithBatcher(traceExporter),
// 		// Document information about this application in resources
// 		sdktrace.WithResource(res),
// 		sdktrace.WithSpanProcessor(bsp),
// 	)
// 	otel.SetTracerProvider(tracerProvider)

// 	// Set global propagator to tracecontext (the default is no-op).
// 	otel.SetTextMapPropagator(propagation.TraceContext{})

// 	// Shutdown will flush any remaining spans and shut down the exporter.
// 	return tracerProvider.Shutdown, nil
// }
