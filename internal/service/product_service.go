package service

import (
	"github.com/pkg/errors"
	"synapsis-challenge/internal/api/request_model"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/entities"
	"synapsis-challenge/internal/repositories"
)

type ProductService interface {
	FetchAllProduct(param request_model.GetProducts) (*[]entities.Product, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (p productService) FetchAllProduct(param request_model.GetProducts) (*[]entities.Product, error) {
	//just simple fetch data from db nothing special
	productsData, err := p.productRepository.FindAllProducts(param)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	return productsData, nil
}
