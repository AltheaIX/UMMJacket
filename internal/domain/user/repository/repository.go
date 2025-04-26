package repository

import "github.com/jmoiron/sqlx"

type UserRepository interface {
	IUserRepository
}

var (
	TableName = "users"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
