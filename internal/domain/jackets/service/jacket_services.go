package service

import (
	"context"
	"github.com/AltheaIX/UMMJacket/internal/domain/jackets/model"
	"github.com/AltheaIX/UMMJacket/internal/domain/jackets/model/dto"
	"github.com/AltheaIX/UMMJacket/shared/filter"
)

type IJacketServices interface {
	InsertJacketsServices(ctx context.Context, request dto.CreateJacketsRequest) (int64, error)
	ResolveJacketsServices(ctx context.Context, filters *filter.Filters) ([]model.Jacket, error)
	UpdateJacketsServices(ctx context.Context, request dto.UpdateJacketsRequest, id int) (int64, error)
	DeleteJacketsServices(ctx context.Context, id int) (int64, error)
}

func (s *JacketsServicesImpl) InsertJacketsServices(ctx context.Context, request dto.CreateJacketsRequest) (
	int64,
	error,
) {
	lastID, err := s.jacketsRepo.InsertJacketsRepository(
		ctx,
		request.Name,
		request.PhotoSizeChart,
		request.PhotoFrontJacket,
		request.PhotoBackJacket,
		request.BasePrice,
		request.ExtraPrice,
	)

	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (s *JacketsServicesImpl) ResolveJacketsServices(ctx context.Context, filters *filter.Filters) (
	[]model.Jacket,
	error,
) {
	jackets, err := s.jacketsRepo.ResolveJacketsRepository(ctx, filters)
	if err != nil {
		return nil, err
	}

	return jackets, nil
}

func (s *JacketsServicesImpl) UpdateJacketsServices(
	ctx context.Context,
	request dto.UpdateJacketsRequest,
	id int,
) (int64, error) {

	rowsAffected, err := s.jacketsRepo.UpdateJacketsRepository(ctx, request, id)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (s *JacketsServicesImpl) DeleteJacketsServices(ctx context.Context, id int) (int64, error) {
	rowsAffected, err := s.jacketsRepo.DeleteJacketsRepository(
		ctx,
		id,
	)

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
