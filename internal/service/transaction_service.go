package service

import (
	"github.com/pkg/errors"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/entities"
	"synapsis-challenge/internal/repositories"
)

type TransactionsService interface {
	FetchAllTransactions(currentUserId string) (*[]entities.Transaction, error)
}

type transactionsService struct {
	transactionsRepository repositories.TransactionsRepository
}

func NewTransactionsService(transactionsRepository repositories.TransactionsRepository) TransactionsService {
	return &transactionsService{
		transactionsRepository: transactionsRepository,
	}
}

func (p transactionsService) FetchAllTransactions(currentUserId string) (*[]entities.Transaction, error) {
	productsData, err := p.transactionsRepository.FetchTransactions(currentUserId)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	return productsData, nil
}
