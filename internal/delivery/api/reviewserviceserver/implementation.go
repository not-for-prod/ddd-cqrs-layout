package reviewserviceserver

import reviewservice "yelp/internal/application/service/review"

type Implementation struct {
	svc *reviewservice.Service
}

func NewImplementation(svc *reviewservice.Service) *Implementation {
	return &Implementation{svc: svc}
}
