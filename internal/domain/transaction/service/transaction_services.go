package service

import (
	"context"
	"github.com/AltheaIX/UMMJacket/internal/domain/transaction/model/dto"
)

type ITransactionServices interface {
	NewTransactionServices(ctx context.Context, userId int, request dto.NewTransactionRequest) (int64, error)
}

func (s *TransactionServicesImpl) NewTransactionServices(
	ctx context.Context,
	userId int,
	request dto.NewTransactionRequest,
) (int64, error) {
	lastId, err := s.transactionRepo.InsertTransactionRepository(ctx, userId, request.ProductId)
	if err != nil {
		return 0, err
	}

	return lastId, nil
}
