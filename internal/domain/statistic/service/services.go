package service

import StatisticRepo "github.com/AltheaIX/UMMJacket/internal/domain/statistic/repository"

type StatisticServices interface {
	IDashboardServices
}

type StatisticServicesImpl struct {
	statisticRepo StatisticRepo.StatisticRepository
}

func NewStatisticServices(statisticRepo StatisticRepo.StatisticRepository) StatisticServices {
	return &StatisticServicesImpl{statisticRepo: statisticRepo}
}
