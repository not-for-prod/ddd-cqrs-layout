package reviewserviceserver

import (
	"context"

	reviewv1 "yelp/internal/generated/pb/yelp/review/v1"
)

func (i *Implementation) ListReviews(
	_ context.Context,
	_ *reviewv1.ListReviewsRequest,
) (*reviewv1.ListReviewsResponse, error) {
	panic("implement me")
}
