package service

import (
	"github.com/pkg/errors"
	"synapsis-challenge/internal/api/request_model"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/entities"
	"synapsis-challenge/internal/repositories"
)

type TransactionsService interface {
	FetchAllTransactions(currentUserId string) (*[]entities.Transaction, error)
	PaymentTransaction(currentUserId string, param request_model.PaymentTransaction) (*entities.Transaction, error)
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
	//fetching all transactions data from db
	productsData, err := p.transactionsRepository.FetchTransactions(currentUserId)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	return productsData, nil
}

func (p transactionsService) PaymentTransaction(currentUserId string, param request_model.PaymentTransaction) (*entities.Transaction, error) {

	//because the deadline is near I only made the simplest payment logic
	//
	//this logic only check amount from request data if amount less than the sub total then throw response error

	transactionData, err := p.transactionsRepository.FetchOneTransactions(currentUserId, param.TransactionId)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	//check if amount less than the total of transactions
	if param.Amount < transactionData.Total {
		return nil, errors.New(consts.InsufficientAmount)
	}

	transactionData.Status = consts.Completed

	//update transactions status
	transactionData, err = p.transactionsRepository.UpdateTransaction(transactionData)

	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	return transactionData, nil
}
