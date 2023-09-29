package repositories

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/entities"
)

type TransactionsRepository interface {
	NewTransactions(transaction *entities.Transaction) (*entities.Transaction, error)
	FetchTransactions(currentUserId string) (*[]entities.Transaction, error)
	FetchOneTransactions(currentUserId, transactionId string) (*entities.Transaction, error)
	UpdateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionsRepo(db *gorm.DB) TransactionsRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) FetchTransactions(currentUserId string) (*[]entities.Transaction, error) {
	var transactions []entities.Transaction

	result := r.db.Preload("Order").Where("transactions.user_id = ?", currentUserId).Order("created_at desc").Find(&transactions)

	return &transactions, result.Error
}

func (r *transactionRepository) NewTransactions(transaction *entities.Transaction) (*entities.Transaction, error) {
	result := r.db.Create(&transaction)

	return transaction, result.Error
}

func (r *transactionRepository) FetchOneTransactions(currentUserId, transactionId string) (*entities.Transaction, error) {
	var transactions entities.Transaction

	result := r.db.Preload("Order").Where("transactions.user_id = ?", currentUserId).Where("transactions.id = ?", transactionId).Order("created_at desc").First(&transactions)

	return &transactions, result.Error
}

func (r *transactionRepository) UpdateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	result := r.db.Model(&transaction).Where("user_id", transaction.UserId).Update("status", transaction.Status).Update("updated_at", transaction.UpdatedAt)
	return transaction, result.Error
}
