package service

import (
	"context"
	"github.com/AltheaIX/UMMJacket/internal/domain/user/model"
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"github.com/rs/zerolog/log"
)

type IUserServices interface {
	GetUsersService(ctx context.Context, filters *filter.Filters) ([]model.User, error)
	GetUserService(ctx context.Context, filters *filter.Filters) (*model.User, error)
}

func (r *UserServiceImpl) GetUsersService(ctx context.Context, filters *filter.Filters) ([]model.User, error) {
	data, err := r.userRepo.ResolveUsersRepository(ctx, filters)
	if err != nil {
		log.Error().Err(err).Msg("[GetUsersService][ResolveUsersRepository] ")
		return nil, err
	}

	return data, nil
}

func (r *UserServiceImpl) GetUserService(ctx context.Context, filters *filter.Filters) (*model.User, error) {
	data, err := r.userRepo.ResolveUsersRepository(ctx, filters)
	if err != nil {
		log.Error().Err(err).Msg("[GetUserService][ResolveUsersRepository] ")
		return nil, err
	}

	var user *model.User

	if len(data) > 0 {
		user = &data[0]
	}

	return user, nil
}
