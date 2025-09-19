package yelp

import (
	"yelp/internal/application/service/yelp/command/add-review"
	"yelp/internal/application/service/yelp/query/list-reviews"
)

type Service struct {
	AddReview  *add_review.Handler
	ListReview *list_reviews.Handler
}

func New() *Service {
	return &Service{}
}
