package repository

import (
	"context"
	"github.com/AltheaIX/UMMJacket/internal/domain/jackets/model"
	"github.com/AltheaIX/UMMJacket/internal/domain/jackets/model/dto"
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"github.com/AltheaIX/UMMJacket/shared/query"
)

type IJacketsRepository interface {
	InsertJacketsRepository(ctx context.Context, args ...interface{}) (int64, error)
	ResolveJacketsRepository(ctx context.Context, filters *filter.Filters) ([]model.Jacket, error)
	UpdateJacketsRepository(ctx context.Context, request dto.UpdateJacketsRequest, id int) (int64, error)
	DeleteJacketsRepository(ctx context.Context, id int) (int64, error)
}

func (r *JacketsRepositoryImpl) InsertJacketsRepository(ctx context.Context, args ...interface{}) (
	int64,
	error,
) {
	result, err := r.exec(
		ctx,
		"INSERT INTO jackets (name, photo_size_chart, photo_front_jacket, photo_back_jacket, base_price, extra_price) VALUES (?, ?, ?, ?, ?, ?)",
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

func (r *JacketsRepositoryImpl) ResolveJacketsRepository(ctx context.Context, filters *filter.Filters) (
	[]model.Jacket,
	error,
) {
	var jackets []model.Jacket

	query, params, err := filter.BuildFilter(filters, TableName)

	err = r.db.SelectContext(ctx, &jackets, query, params...)
	if err != nil {
		return nil, err
	}

	return jackets, nil
}

func (r *JacketsRepositoryImpl) UpdateJacketsRepository(ctx context.Context, request dto.UpdateJacketsRequest, id int) (
	int64,
	error,
) {
	q, p, err := query.BuildUpdateQuery(request, TableName, id)
	if err != nil {
		return 0, err
	}

	result, err := r.exec(
		ctx,
		q,
		p...,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *JacketsRepositoryImpl) DeleteJacketsRepository(ctx context.Context, id int) (int64, error) {
	result, err := r.exec(
		ctx,
		"DELETE FROM jackets WHERE id=?",
		id,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
