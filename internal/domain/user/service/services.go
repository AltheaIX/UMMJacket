package service

import (
	UserRepository "github.com/AltheaIX/UMMJacket/internal/domain/user/repository"
)

type UserServices interface {
	IUserServices
}

type UserServiceImpl struct {
	userRepo UserRepository.UserRepository
}

func NewUserService(userRepo UserRepository.UserRepository) UserServices {
	return &UserServiceImpl{userRepo: userRepo}
}
