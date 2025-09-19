package review_service_server

import (
	"context"

	add_review "yelp/internal/application/service/yelp/command/add-review"
	reviewv1 "yelp/internal/generated/pb/yelp/common/review/v1"
	commonv1 "yelp/internal/generated/pb/yelp/common/v1"
	reviewSvcv1 "yelp/internal/generated/pb/yelp/review/v1"
	"yelp/internal/pkg/prospan"
)

func (i *Implementation) AddReview(
	ctx context.Context,
	req *reviewSvcv1.AddReviewRequest,
) (*reviewSvcv1.AddReviewResponse, error) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	result, err := i.yelpService.AddReview.Execute(
		ctx, add_review.Command{
			ClientID:    "",
			Title:       "",
			Description: "",
		},
	)
	if err != nil {
		return nil, err
	}

	return &reviewSvcv1.AddReviewResponse{
		Id: &reviewv1.ID{Value: &commonv1.UUID{Value: result.ID.String()}},
	}, nil
}
