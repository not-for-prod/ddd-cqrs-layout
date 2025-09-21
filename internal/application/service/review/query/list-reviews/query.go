package listreviews

import (
	"context"

	applicationdto "yelp/internal/application/service/review/dto"
	"yelp/internal/domain/entity/review"
	"yelp/internal/pkg/prospan"
)

type Query struct {
	Limit  int
	Offset int
}

type Result struct {
	Reviews []review.Review
	HasMore bool
}

type Handler struct {
	reviewRepository reviewRepository
}

func NewCommandHandler(reviewRepository reviewRepository) *Handler {
	return &Handler{
		reviewRepository: reviewRepository,
	}
}

func (h *Handler) Execute(ctx context.Context, q Query) (Result, error) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	list, err := h.reviewRepository.List(
		ctx, applicationdto.ListReviewRequest{
			Limit:  q.Limit,
			Offset: q.Offset,
		},
	)
	if err != nil {
		return Result{}, err
	}

	return Result{
		Reviews: list.Reviews,
		HasMore: list.HasMore,
	}, nil
}
