package repository

import "context"

type ITransactionRepository interface {
	InsertTransactionRepository(ctx context.Context, args ...interface{}) (int64, error)
}

func (r *TransactionRepositoryImpl) InsertTransactionRepository(ctx context.Context, args ...interface{}) (
	int64,
	error,
) {
	result, err := r.exec(
		ctx,
		"INSERT INTO transactions (user_id, product_id) VALUES (?, ?)",
		args...,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
