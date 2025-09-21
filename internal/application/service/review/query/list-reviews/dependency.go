package listreviews

import (
	"context"

	applicationdto "yelp/internal/application/service/review/dto"
)

//go:generate moq -out mock.go . reviewRepository

type reviewRepository interface {
	List(ctx context.Context, req applicationdto.ListReviewRequest) (applicationdto.ListReviewResponse, error)
}
