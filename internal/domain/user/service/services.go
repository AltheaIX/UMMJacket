package service

import (
	UserRepository "github.com/AltheaIX/UMMJacket/internal/domain/user/repository"
)

type UserServices interface {
	IUserServices
}

type UserServicesImpl struct {
	userRepo UserRepository.UserRepository
}

func NewUserServices(userRepo UserRepository.UserRepository) UserServices {
	return &UserServicesImpl{userRepo: userRepo}
}
