package addreview

import (
	"context"

	"yelp/internal/domain/entity/client"
	review "yelp/internal/domain/entity/review"
	"yelp/internal/pkg/prospan"

	"golang.org/x/sync/errgroup"
)

type Command struct {
	ClientID    client.ID
	Title       string
	Description string
}

type Result struct {
	ID review.ID
}

type Handler struct {
	reviewRepository reviewRepository
	clientRepository clientRepository
	txManager        txManager
}

func NewCommandHandler(
	reviewRepository reviewRepository,
	clientRepository clientRepository,
	txManager txManager,
) *Handler {
	return &Handler{
		reviewRepository: reviewRepository,
		clientRepository: clientRepository,
		txManager:        txManager,
	}
}

func (h *Handler) Execute(ctx context.Context, cmd Command) (Result, error) {
	ctx, span := prospan.Start(ctx)
	defer span.End()

	group, ctx := errgroup.WithContext(ctx)
	review := review.New(cmd.ClientID, cmd.Title, cmd.Description)

	err := h.txManager.Do(
		ctx, func(ctx context.Context) error {
			group.Go(func() error { return h.reviewRepository.Add(ctx, *review) })
			group.Go(func() error { return h.clientRepository.IncReviewCount(ctx, cmd.ClientID) })

			return group.Wait()
		},
	)
	if err != nil {
		return Result{}, err
	}

	return Result{
		ID: review.ID,
	}, nil
}
