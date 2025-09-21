package applicationdto

import "yelp/internal/domain/entity/review"

type ListReviewRequest struct {
	Limit, Offset int
}

type ListReviewResponse struct {
	Reviews []review.Review
	HasMore bool
}
