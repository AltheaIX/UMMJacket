package repository

import (
	"context"
	"github.com/AltheaIX/UMMJacket/internal/domain/user/model"
	"github.com/AltheaIX/UMMJacket/shared/filter"
)

type IUserRepository interface {
	ResolveUsersRepository(ctx context.Context, filters *filter.Filters) ([]model.User, error)
}

func (r *UserRepositoryImpl) ResolveUsersRepository(ctx context.Context, filters *filter.Filters) ([]model.User, error) {
	var users []model.User

	query, params, err := filter.BuildFilterAnd(filters.Filter, TableName)

	err = r.db.SelectContext(ctx, &users, query, params...)
	if err != nil {
		return nil, err
	}

	return users, nil
}
