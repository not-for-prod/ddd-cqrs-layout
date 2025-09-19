package model

import (
	"yelp/internal/domain/entity/client"
	"yelp/internal/domain/entity/review"
)

type Review struct {
	ClientID    string `db:"client_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

func ReviewFromDomain(r review.Review) Review {
	return Review{
		ClientID:    string(r.ClientID),
		Title:       r.Title,
		Description: r.Description,
	}
}

func (r *Review) ToDomain() review.Review {
	return review.Review{
		ClientID:    client.ID(r.ClientID),
		Title:       r.Title,
		Description: r.Description,
	}
}
