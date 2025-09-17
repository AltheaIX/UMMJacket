package service

import JacketsRepository "github.com/AltheaIX/UMMJacket/internal/domain/jackets/repository"

type JacketsServices interface {
	IJacketServices
}

type JacketsServicesImpl struct {
	jacketsRepo JacketsRepository.JacketsRepository
}

func NewJacketsServices(jacketsRepo JacketsRepository.JacketsRepository) JacketsServices {
	return &JacketsServicesImpl{
		jacketsRepo: jacketsRepo,
	}
}
