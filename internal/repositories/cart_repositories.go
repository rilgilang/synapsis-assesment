package repositories

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/entities"
	"time"
)

type CartRepository interface {
	CreateCart(cart *entities.Cart) error
	CheckCart(userId string) (*entities.Cart, error)
	UpdateCart(cart *entities.Cart) (*entities.Cart, error)
	FetchCart(userId string) (*entities.Cart, error)
}
type cartRepository struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) CreateCart(cart *entities.Cart) error {
	err := r.db.Create(cart).Error
	return err
}

func (r *cartRepository) CheckCart(userId string) (*entities.Cart, error) {
	cart := entities.Cart{}
	err := r.db.Where("user_id = ?", userId).First(&cart).Error

	if err != nil && err.Error() == consts.SqlNoRow {
		return nil, nil
	}

	return &cart, err
}

func (r *cartRepository) UpdateCart(cart *entities.Cart) (*entities.Cart, error) {
	cart.UpdatedAt = time.Now()
	result := r.db.Model(&cart).Update("total", cart.Total).Update("updated_at", cart.UpdatedAt).Where("user_id", cart.UserId)
	return cart, result.Error
}

func (r *cartRepository) FetchCart(userId string) (*entities.Cart, error) {
	cart := entities.Cart{}
	err := r.db.Preload("CartProduct").Where("user_id = ?", userId).First(&cart).Error

	if err != nil && err.Error() == consts.SqlNoRow {
		return nil, nil
	}

	return &cart, err
}
