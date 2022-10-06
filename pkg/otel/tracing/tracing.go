package tracing

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/core"
	otel2 "github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel"
	"go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel"
	_ "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"os"
)

//https://opentelemetry.io/docs/reference/specification/
//https://opentelemetry.io/docs/instrumentation/go/getting-started/
//https://opentelemetry.io/docs/instrumentation/go/manual/
//https://opentelemetry.io/docs/reference/specification/trace/semantic_conventions/
//https://uptrace.dev/opentelemetry/go-tracing.html
//https://lightstep.com/blog/opentelemetry-go-all-you-need-to-know
//https://trstringer.com/otel-part2-instrumentation/
//https://trstringer.com/otel-part5-propagation/
//https://github.com/tedsuo/otel-go-basics/blob/main/server.go

type openTelemtry struct {
	config         *otel2.OpenTelemetryConfig
	tracerProvider *tracesdk.TracerProvider
	jaegerExporter tracesdk.SpanExporter
	zipkinExporter tracesdk.SpanExporter
	stdExporter    tracesdk.SpanExporter
	tracerName     string
}

// Create one tracer per package
// NOTE: You only need a tracer if you are creating your own spans

// Tracer global tracer for app
var Tracer trace.Tracer

func init() {
	Tracer = NewCustomTracer("app-tracer") //instrumentation name
}

func AddOtelTracing(config *otel2.OpenTelemetryConfig) (*tracesdk.TracerProvider, error) {
	openTel := &openTelemtry{config: config}

	err := openTel.configExporters()
	if err != nil {
		return nil, errors.WrapIf(err, "error in config exporter")
	}

	//https://opentelemetry.io/docs/instrumentation/go/manual/#initializing-a-new-tracer
	err = openTel.configTracerProvider()
	if err != nil {
		return nil, err
	}

	//https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/propagators/ot/ot_propagator.go
	//https://github.com/open-telemetry/opentelemetry-go/blob/main/propagation/trace_context.go
	//https://github.com/open-telemetry/opentelemetry-go/blob/main/propagation/baggage.go/
	//https://trstringer.com/otel-part5-propagation/
	propagators := []propagation.TextMapPropagator{
		ot.OT{}, // should be placed before `TraceContext` for preventing conflict
		propagation.TraceContext{},
		propagation.Baggage{},
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(openTel.tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagators...))

	//https://trstringer.com/otel-part2-instrumentation/
	// Finally, set the tracer that can be used for this package.
	Tracer = NewCustomTracer(config.InstrumetationName)

	return openTel.tracerProvider, nil
}

func (o *openTelemtry) configTracerProvider() error {
	var sampler tracesdk.Sampler
	if o.config.AlwaysOnSampler {
		sampler = tracesdk.AlwaysSample()
	} else {
		sampler = tracesdk.NeverSample()
	}

	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(o.config.ServiceName),
			attribute.Int64("ID", o.config.Id),
			attribute.String("environment", core.GetEnvironment()),
		),
	)
	if err != nil {
		return err
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(o.jaegerExporter),
		tracesdk.WithBatcher(o.zipkinExporter),
		tracesdk.WithBatcher(o.stdExporter),
		tracesdk.WithSampler(sampler),

		// https://opentelemetry.io/docs/instrumentation/go/exporting_data/#resources
		//Resources are a special type of attribute that apply to all spans generated by a process
		tracesdk.WithResource(r),
	)
	o.tracerProvider = tp
	return nil
}

func (o *openTelemtry) configExporters() error {
	logger := log.New(os.Stderr, "otel_log", log.Ldate|log.Ltime|log.Llongfile)

	if o.config.JaegerExporterConfig != nil {
		// Create the Jaeger exporter
		jaegerExporter, err := jaeger.New(jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost(o.config.JaegerExporterConfig.AgentHost),
			jaeger.WithAgentPort(o.config.JaegerExporterConfig.AgentPort),
			jaeger.WithLogger(logger),
		))
		if err != nil {
			return err
		}
		o.jaegerExporter = jaegerExporter
	}
	if o.config.ZipkinExporterConfig != nil {
		zipkinExporter, err := zipkin.New(
			o.config.ZipkinExporterConfig.Url,
		)
		if err != nil {
			return err
		}

		o.zipkinExporter = zipkinExporter
	}
	if o.config.UseStdout {
		stdExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
		if err != nil {
			return fmt.Errorf("creating stdout exporter: %w", err)
		}
		o.stdExporter = stdExporter
	}

	return nil
}
