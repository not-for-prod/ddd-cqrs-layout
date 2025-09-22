package internal

import (
	"context"

	txManager "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func initLogger(lc fx.Lifecycle) {
	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				return nil
			},
			OnStop: func(_ context.Context) error {
				return nil
			},
		},
	)
}

func initTracer(lc fx.Lifecycle) {
	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				return nil
			},
			OnStop: func(_ context.Context) error {
				return nil
			},
		},
	)
}

func initDB(lc fx.Lifecycle) *sqlx.DB {
	lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				return nil
			},
			OnStop: func(_ context.Context) error {
				return nil
			},
		},
	)

	return nil
}

func initTxManager(db *sqlx.DB) *manager.Manager {
	return manager.Must(txManager.NewDefaultFactory(db))
}
