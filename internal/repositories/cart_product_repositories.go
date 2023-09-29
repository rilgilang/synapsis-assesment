package repositories

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/entities"
	"time"
)

type CartProductRepository interface {
	ProductInsertToCart(product *entities.CartProduct) error
	FindProductInCart(userId, cartID string) (*[]entities.CartProduct, error)
	FindOneProductInCart(userId, productCartId, cartID string) (*entities.CartProduct, error)
	UpdateProductInCart(cartId, productCartId string, quantity, total int) (*entities.CartProduct, error)
	DeleteProductFromCart(cartId, productCartId, userId string) error
	DeleteAllProductFromCart(cartId, userId string) error
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

func (r *cartProductRepository) FindOneProductInCart(userId, productCartId, cartId string) (*entities.CartProduct, error) {
	var products entities.CartProduct

	result := r.db.Where("user_id = ?", userId).Where("id = ?", productCartId).Where("cart_id = ?", cartId).Where("deleted_at IS NULL").First(&products)
	if result.Error != nil && result.Error.Error() == consts.SqlNoRow {
		return nil, nil
	}

	return &products, result.Error
}

func (r *cartProductRepository) UpdateProductInCart(cartId, productCartId string, quantity, total int) (*entities.CartProduct, error) {
	cartProduct := entities.CartProduct{}
	result := r.db.Model(&cartProduct).Where("id", productCartId).Where("cart_id", cartId).Update("quantity", quantity).Update("total", total).Update("updated_at", time.Now())

	return &cartProduct, result.Error
}

func (r *cartProductRepository) DeleteProductFromCart(cartId, productCartId, userId string) error {
	err := r.db.Where("id = ?", productCartId).Where("cart_id = ?", cartId).Where("user_id = ?", userId).Delete(&entities.CartProduct{}).Error

	return err
}

func (r *cartProductRepository) DeleteAllProductFromCart(cartId, userId string) error {
	err := r.db.Where("cart_id = ?", cartId).Where("user_id = ?", userId).Delete(&entities.CartProduct{}).Error

	return err
}
