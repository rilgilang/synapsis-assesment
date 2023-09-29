package repositories

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/entities"
)

// Repository interface allows us to access the CRUD Operations in sql here.
type TransactionsRepository interface {
	NewTransactions(transaction *entities.Transaction) (*entities.Transaction, error)
	FetchTransactions(currentUserId string) (*[]entities.Transaction, error)
}
type transactionRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewTransactionsRepo(db *gorm.DB) TransactionsRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r transactionRepository) FetchTransactions(currentUserId string) (*[]entities.Transaction, error) {
	var transactions []entities.Transaction

	result := r.db.Preload("Order").Where("transactions.user_id = ?", currentUserId).Order("created_at desc").Find(&transactions)

	return &transactions, result.Error
}

func (r transactionRepository) NewTransactions(transaction *entities.Transaction) (*entities.Transaction, error) {
	result := r.db.Create(&transaction)

	return transaction, result.Error
}
