package service

import TransactionRepository "github.com/AltheaIX/UMMJacket/internal/domain/transaction/repository"

type TransactionServices interface {
	ITransactionServices
}

type TransactionServicesImpl struct {
	transactionRepo TransactionRepository.TransactionRepository
}

func NewTransactionServices(transactionRepo TransactionRepository.TransactionRepository) TransactionServices {
	return &TransactionServicesImpl{transactionRepo: transactionRepo}
}
