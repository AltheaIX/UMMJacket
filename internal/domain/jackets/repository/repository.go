package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type JacketsRepository interface {
	IJacketsRepository
}

var TableName = "jackets"

type JacketsRepositoryImpl struct {
	db *sqlx.DB
}

func NewJacketRepository(db *sqlx.DB) JacketsRepository {
	return &JacketsRepositoryImpl{
		db: db,
	}
}

func (r *JacketsRepositoryImpl) exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
