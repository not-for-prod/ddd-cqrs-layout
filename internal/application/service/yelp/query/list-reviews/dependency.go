package list_reviews

import "yelp/internal/domain/entity/review"

type reviewRepository interface {
	List(ctx context.Context, clientID string) ([]review.Review, error)
}
