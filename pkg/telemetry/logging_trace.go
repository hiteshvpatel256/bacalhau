package telemetry

import (
	"context"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/baggage"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

type loggingTracerProvider struct {
	embedded.TracerProvider
	delegate shutdownTracerProvider
}

func (l loggingTracerProvider) Shutdown(ctx context.Context) error {
	return l.delegate.Shutdown(ctx)
}

func (l loggingTracerProvider) Tracer(name string, options ...oteltrace.TracerOption) oteltrace.Tracer {
	tracer := l.delegate.Tracer(name, options...)
	return loggingTracer{
		delegate: tracer,
	}
}

type loggingTracer struct {
	embedded.Tracer
	delegate oteltrace.Tracer
}

func (l loggingTracer) Start(ctx context.Context, spanName string, opts ...oteltrace.SpanStartOption) (context.Context, oteltrace.Span) {
	ctx, span := l.delegate.Start(ctx, spanName, opts...)

	// Name of the attributes taken from https://opentelemetry.io/docs/reference/specification/logs/#json-formats
	logger := zerolog.Ctx(ctx).With().
		Stringer("trace_id", span.SpanContext().TraceID()).
		Stringer("span_id", span.SpanContext().SpanID())

	if v := baggage.FromContext(ctx).Member(TracerAttributeNameJobID).Value(); v != "" {
		logger = logger.Str(TracerAttributeNameJobID, v)
	}

	ctx = logger.Logger().WithContext(ctx)

	return ctx, span
}

var _ shutdownTracerProvider = loggingTracerProvider{}
var _ oteltrace.Tracer = loggingTracer{}
