package review

import (
	"context"

	"github.com/samber/lo"
	"yelp/internal/application/service/yelp/dto"
	"yelp/internal/domain/entity/review"
	"yelp/internal/infrastructure/repository/review/model"
	"yelp/internal/infrastructure/repository/review/query"
	"yelp/internal/pkg/prospan"
)

func (r *Repository) List(
	ctx context.Context,
	req applicationdto.ListReviewRequest,
) (applicationdto.ListReviewResponse, error) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	var reviews []model.Review

	err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).SelectContext(ctx, &reviews, query.List, req.Limit, req.Offset)
	if err != nil {
		span.RecordError(err)
		return applicationdto.ListReviewResponse{}, err
	}

	return applicationdto.ListReviewResponse{
		Reviews: lo.Map(
			reviews, func(item model.Review, index int) review.Review {
				return item.ToDomain()
			},
		),
	}, nil
}
