package expirationworker

import (
	"context"
	"time"

	"yelp/internal/config"

	"golang.org/x/sync/errgroup"
)

type Worker struct {
	reviewRepository reviewRepository
}

func New(reviewRepository reviewRepository) *Worker {
	return &Worker{
		reviewRepository: reviewRepository,
	}
}

func (w *Worker) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	for range config.GetInstance().ExpirationWorker.Count {
		group.Go(
			func() error {
				w.work(ctx)
				return nil
			},
		)
	}

	return group.Wait()
}

func (w *Worker) work(ctx context.Context) {
	timer := time.NewTicker(config.GetInstance().ExpirationWorker.Interval)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			w.tick()
		case <-ctx.Done():
			return
		}
	}
}

func (w *Worker) tick() {
	// actual work
}
