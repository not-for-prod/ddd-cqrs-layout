package internal

import (
	"context"
	"log/slog"

	expirationworker "yelp/internal/application/woker/expiration"
	review_service_server "yelp/internal/delivery/api/reviewserviceserver"
	reviewv1 "yelp/internal/generated/pb/yelp/review/v1"

	"github.com/not-for-prod/clay/server"
	"go.uber.org/fx"
)

func runYelpReviewService(lc fx.Lifecycle, shutdowner fx.Shutdowner, svc *review_service_server.Implementation) {
	serviceServer := server.NewServer(12345)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					err := serviceServer.Run(reviewv1.NewReviewServiceServiceDesc(svc))
					if err != nil {
						_ = shutdowner.Shutdown(fx.ExitCode(1))
					}

					return
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return serviceServer.Stop(ctx)
			},
		},
	)
}

func runExpirationWorker(lc fx.Lifecycle) {
	workerPool := expirationworker.New()
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Create a cancellable context for the worker
				workerCtx, cancel := context.WithCancel(ctx)

				// Store cancel in OnStop
				lc.Append(
					fx.Hook{
						OnStop: func(context.Context) error {
							cancel()
							return nil
						},
					},
				)

				go func() {
					if err := workerPool.Run(workerCtx); err != nil {
						slog.Error("expiration worker exited with error", err)
					}
				}()

				return nil
			},
		},
	)
}
