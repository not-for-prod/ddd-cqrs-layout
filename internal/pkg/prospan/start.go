package prospan

import (
	"context"

	"yelp/internal/pkg/prospan/autoname"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const initialSkipFrames = 2

type ProSpan struct {
	trace.Span
}

//nolint:spancheck // its others
func Start(ctx context.Context) (context.Context, ProSpan) {
	ctx, span := otel.Tracer("").Start(ctx, autoname.GetRuntimeFunc(initialSkipFrames))
	return ctx, ProSpan{span}
}

func (s *ProSpan) End(options ...trace.SpanEndOption) {
	s.Span.End(options...)
}

func (s *ProSpan) RecordError(err error) {
	s.Span.RecordError(err)
}
