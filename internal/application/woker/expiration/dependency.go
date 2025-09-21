package expirationworker

import (
	"context"

	"yelp/internal/domain/entity/review"
)

//go:generate moq -out mock.go . reviewRepository

type reviewRepository interface {
	List(ctx context.Context) ([]review.Review, error)
}
