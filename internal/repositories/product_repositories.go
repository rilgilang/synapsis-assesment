package repositories

import (
	"gorm.io/gorm"
	"synapsis-challenge/internal/api/request_model"
	"synapsis-challenge/internal/entities"
)

// Repository interface allows us to access the CRUD Operations in sql here.
type ProductRepository interface {
	FindAllProducts(param request_model.GetProducts) (*[]entities.Product, error)
	FindOneProducts(productId string) (*entities.Product, error)
}
type productRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewProductRepo(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p productRepository) FindAllProducts(param request_model.GetProducts) (*[]entities.Product, error) {
	var products []entities.Product

	result := p.db

	if param.Category != "" {
		result = p.db.Preload("ProductCategory").Joins("LEFT JOIN product_categories on product_categories.id = products.category_id").Where("product_categories.category_name = ?", param.Category).Where("products.deleted_at IS NULL").Order("created_at desc").Find(&products)
	} else {
		result = p.db.Preload("ProductCategory").Joins("LEFT JOIN product_categories on product_categories.id = products.category_id").Where("products.deleted_at IS NULL").Order("created_at desc").Find(&products)
	}

	return &products, result.Error
}

func (p productRepository) FindOneProducts(productId string) (*entities.Product, error) {
	var products entities.Product

	result := p.db.Preload("ProductCategory").Joins("LEFT JOIN product_categories on product_categories.id = products.category_id").Where("products.id = ?", productId).Where("products.deleted_at IS NULL").Order("created_at desc").First(&products)

	return &products, result.Error
}
