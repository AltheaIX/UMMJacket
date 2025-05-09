package repository

import "github.com/jmoiron/sqlx"

type StatisticRepository interface {
	IDashboardRepository
}

type StatisticRepositoryImpl struct {
	db *sqlx.DB
}

func NewStatisticRepository(db *sqlx.DB) StatisticRepository {
	return &StatisticRepositoryImpl{db: db}
}
