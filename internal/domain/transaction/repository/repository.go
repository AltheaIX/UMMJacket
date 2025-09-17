package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	ITransactionRepository

	exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type TransactionRepositoryImpl struct {
	db *sqlx.DB
}

func NewTransactionRepositoryImpl(db *sqlx.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
