package list_reviews

import (
	"context"

	"yelp/internal/pkg/prospan"
)

type Query struct{}

type Result struct{}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Execute(ctx context.Context, q Query) (Result, error) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	return Result{}, nil
}
