package review_service_server

import "yelp/internal/application/service/yelp"

type Implementation struct {
	yelpService *yelp.Service
}

func NewImplementation(service *yelp.Service) *Implementation {
	return &Implementation{
		yelpService: service,
	}
}
