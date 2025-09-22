package internal

import (
	txManager "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func initLogger(_ fx.Lifecycle) {
	panic("not implemented")
}

func initTracer(_ fx.Lifecycle) {
	panic("not implemented")
}

func initDB(_ fx.Lifecycle) *sqlx.DB {
	panic("not implemented")
}

func initTxManager(db *sqlx.DB) *manager.Manager {
	return manager.Must(txManager.NewDefaultFactory(db))
}
