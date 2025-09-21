package internal

import (
	"yelp/internal/application/service/review"
	reviewServiceServer "yelp/internal/delivery/api/reviewserviceserver"
	reviewrepository "yelp/internal/infrastructure/repository/review"

	"go.uber.org/fx"
)

func Run() {
	app := fx.New(
		fx.Invoke(
			initLogger,
			initTracer,
		),
		fx.Provide(
			initDB,
			initTxManager,
			reviewrepository.New,
			reviewservice.New,
			reviewServiceServer.NewImplementation,
		),
		fx.Invoke(runYelpReviewService),
		fx.Invoke(runExpirationWorker),
	)
	app.Run()
}
