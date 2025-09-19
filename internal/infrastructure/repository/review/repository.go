package review

import (
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db        *sqlx.DB
	ctxGetter *trmsqlx.CtxGetter
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		db:        sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx"),
		ctxGetter: trmsqlx.DefaultCtxGetter,
	}
}
