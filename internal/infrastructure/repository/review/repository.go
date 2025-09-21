package reviewrepository

import (
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db        *sqlx.DB
	ctxGetter *trmsqlx.CtxGetter
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db:        db,
		ctxGetter: trmsqlx.DefaultCtxGetter,
	}
}
