package review_service_server

import (
	"context"
	reviewv1 "yelp/internal/generated/pb/yelp/review/v1"
)

func (i *Implementation) ListReviews(arg0 context.Context, arg1 *reviewv1.ListReviewsRequest) (*reviewv1.ListReviewsResponse, error) {
	panic("implement me")
}
