package reviewservice

import (
	addreview "yelp/internal/application/service/review/command/add-review"
	listreviews "yelp/internal/application/service/review/query/list-reviews"
	reviewrepository "yelp/internal/infrastructure/repository/review"

	"github.com/avito-tech/go-transaction-manager/trm/manager"
)

type Service struct {
	AddReview  *addreview.Handler
	ListReview *listreviews.Handler
}

func New(reviewRepository *reviewrepository.Repository, txManager *manager.Manager) *Service {
	return &Service{
		AddReview:  addreview.NewCommandHandler(reviewRepository, nil, txManager),
		ListReview: listreviews.NewCommandHandler(reviewRepository),
	}
}
