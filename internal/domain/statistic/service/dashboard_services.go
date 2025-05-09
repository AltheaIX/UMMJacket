package service

import (
	"context"
	"github.com/AltheaIX/UMMJacket/internal/domain/statistic/model"
	"github.com/rs/zerolog/log"
)

type IDashboardServices interface {
	GetDashboardCount(ctx context.Context) (model.Statistic, error)
}

func (s *StatisticServiceImpl) GetDashboardCount(ctx context.Context) (model.Statistic, error) {
	statistic, err := s.statisticRepo.GetDashboardCounts(ctx)
	if err != nil {
		log.Error().Err(err).Msg("[GetDashboardCount][GetDashboardCounts]")
		return statistic, err
	}

	return statistic, nil
}
