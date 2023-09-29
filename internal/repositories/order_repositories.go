package repositories

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/entities"
)

type OrderRepository interface {
	InsertOrder(order *entities.Order) error
}
type orderRepository struct {
	db *gorm.DB
}

func NewOrdersRepo(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) InsertOrder(order *entities.Order) error {
	result := r.db.Create(&order)

	return result.Error
}
