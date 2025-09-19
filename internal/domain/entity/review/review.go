package review

import (
	"yelp/internal/domain/entity/client"

	"github.com/google/uuid"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Review struct {
	ID          ID
	ClientID    client.ID
	Title       string
	Description string
}

func New(clientID client.ID, title, description string) *Review {
	return &Review{
		ID:          ID(uuid.NewString()),
		ClientID:    clientID,
		Title:       title,
		Description: description,
	}
}
