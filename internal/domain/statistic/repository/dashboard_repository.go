package repository

import (
	"context"
	"github.com/AltheaIX/UMMJacket/internal/domain/statistic/model"
	"github.com/AltheaIX/UMMJacket/shared/filter"
)

type IDashboardRepository interface {
	GetDashboardCounts(ctx context.Context) (model.Statistic, error)
}

func (r *StatisticRepositoryImpl) GetDashboardCounts(ctx context.Context) (model.Statistic, error) {
	var statistic model.Statistic

	var tables = []string{
		"users",
		"transactions",
	}

	query := filter.GetMultipleTableCounts(tables)

	err := r.db.GetContext(ctx, &statistic, query)
	if err != nil {
		return statistic, err
	}

	return statistic, nil
}
