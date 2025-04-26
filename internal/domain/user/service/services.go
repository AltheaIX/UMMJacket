package service

import (
	UserRepository "github.com/AltheaIX/UMMJacket/internal/domain/user/repository"
)

type UserService interface {
	IUserServices
}

type UserServiceImpl struct {
	userRepo UserRepository.UserRepository
}

func NewUserService(userRepo UserRepository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}
