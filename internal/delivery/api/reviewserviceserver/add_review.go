package reviewserviceserver

import (
	"context"

	reviewv1 "yelp/internal/generated/pb/yelp/review/v1"
)

func (i *Implementation) AddReview(
	_ context.Context,
	_ *reviewv1.AddReviewRequest,
) (*reviewv1.AddReviewResponse, error) {
	panic("implement me")
}
