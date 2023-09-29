package repositories

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/entities"
	"time"
)

type CartProductRepository interface {
	ProductInsertToCart(product *entities.CartProduct) error
	FindProductInCart(userId, cartID string) (*[]entities.CartProduct, error)
	UpdateProductInCart(cartId, productCartId string, quantity, total int) (*entities.CartProduct, error)
}
type cartProductRepository struct {
	db *gorm.DB
}

func NewCartProductRepo(db *gorm.DB) CartProductRepository {
	return &cartProductRepository{
		db: db,
	}
}

func (r *cartProductRepository) ProductInsertToCart(product *entities.CartProduct) error {
	err := r.db.Create(product).Error
	return err
}

func (r *cartProductRepository) FindProductInCart(userId, cartId string) (*[]entities.CartProduct, error) {
	var products []entities.CartProduct

	result := r.db.Where("user_id = ?", userId).Where("cart_id = ?", cartId).Where("deleted_at IS NULL").Find(&products)

	return &products, result.Error
}
func (r *cartProductRepository) UpdateProductInCart(cartId, productCartId string, quantity, total int) (*entities.CartProduct, error) {
	cartProduct := entities.CartProduct{}
	result := r.db.Model(&cartProduct).Where("id", productCartId).Where("cart_id", cartId).Update("quantity", quantity).Update("total", total).Update("updated_at", time.Now())

	return &cartProduct, result.Error
}
