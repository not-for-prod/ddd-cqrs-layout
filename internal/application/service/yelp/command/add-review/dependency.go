package add_review

import (
	"context"

	"yelp/internal/domain/entity/client"
	"yelp/internal/domain/entity/review"
)

//go:generate moq -out mock.go . reviewRepository

type reviewRepository interface {
	Add(ctx context.Context, review review.Review) error
}

type clientRepository interface {
	IncReviewCount(ctx context.Context, clientID client.ID) error
}

type txManager interface {
	Do(ctx context.Context, fn func(context.Context) error) error
}
