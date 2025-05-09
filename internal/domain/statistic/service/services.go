package service

import StatisticRepo "github.com/AltheaIX/UMMJacket/internal/domain/statistic/repository"

type StatisticServices interface {
	IDashboardServices
}

type StatisticServiceImpl struct {
	statisticRepo StatisticRepo.StatisticRepository
}

func NewStatisticService(statisticRepo StatisticRepo.StatisticRepository) StatisticServices {
	return &StatisticServiceImpl{statisticRepo: statisticRepo}
}
