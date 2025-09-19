package review

import (
	"context"

	"yelp/internal/domain/entity/review"
	"yelp/internal/infrastructure/repository/review/model"
	"yelp/internal/infrastructure/repository/review/query"
	"yelp/internal/pkg/prospan"
)

func (r *Repository) Add(ctx context.Context, review review.Review) error {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	dto := model.ReviewFromDomain(review)

	_, err := r.ctxGetter.DefaultTrOrDB(ctx, r.db).ExecContext(
		ctx,
		query.Insert,
		dto.ClientID,
		dto.Title,
		dto.Description,
	)
	if err != nil {
		span.RecordError(err)
		return err
	}

	return nil
}
