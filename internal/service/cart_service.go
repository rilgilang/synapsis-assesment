package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"synapsis-challenge/internal/api/request_model"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/entities"
	"synapsis-challenge/internal/repositories"
	"time"
)

type CartService interface {
	AddToCart(currentUserId string, param request_model.AddToCart) (*entities.Cart, error)
}

type cartService struct {
	productRepository     repositories.ProductRepository
	cartRepository        repositories.CartRepository
	productCartRepository repositories.CartProductRepository
}

func NewCartService(
	productRepository repositories.ProductRepository,
	cartRepository repositories.CartRepository,
	productCartRepository repositories.CartProductRepository,
) CartService {
	return &cartService{
		productRepository:     productRepository,
		cartRepository:        cartRepository,
		productCartRepository: productCartRepository,
	}
}

func (p cartService) AddToCart(currentUserId string, param request_model.AddToCart) (*entities.Cart, error) {
	//get the product detail
	productsData, err := p.productRepository.FindOneProducts(param.ProductId)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	//find user cart
	cart, err := p.cartRepository.CheckCart(currentUserId)

	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	//if user has no cart yet then create a new cart
	if cart == nil {
		err = p.cartRepository.CreateCart(&entities.Cart{
			ID:        uuid.New().String(),
			UserId:    currentUserId,
			Total:     0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		})

		if err != nil {
			return nil, errors.New(consts.InternalServerError)
		}

		cart, err = p.cartRepository.CheckCart(currentUserId)

		if err != nil {
			return nil, errors.New(consts.InternalServerError)
		}
	}

	//check if the product already in cart
	productsInCart, err := p.productCartRepository.FindProductInCart(currentUserId, cart.ID)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	checkProduct := map[string]string{}
	checkProduct["alreadyInCart"] = "false"
	newCartProductQuantity := param.Total
	newCartProductTotal := 0
	subTotal := 0

	for _, productCart := range *productsInCart {
		if productCart.ProductId == param.ProductId {
			checkProduct["alreadyInCart"] = "true"
			checkProduct["productCartId"] = productCart.ID

			newCartProductTotal = newCartProductQuantity * productsData.Price

			subTotal += newCartProductTotal
		} else {
			subTotal += productCart.Total
		}
	}

	if checkProduct["alreadyInCart"] == "true" {
		_, err = p.productCartRepository.UpdateProductInCart(cart.ID, checkProduct["productCartId"], newCartProductQuantity, newCartProductTotal)

		if err != nil {
			return nil, errors.New(consts.InternalServerError)
		}

	} else {
		//inserting product to cart
		err = p.productCartRepository.ProductInsertToCart(&entities.CartProduct{
			ID:        uuid.New().String(),
			UserId:    currentUserId,
			CartId:    cart.ID,
			ProductId: param.ProductId,
			Quantity:  param.Total,
			Total:     param.Total * productsData.Price,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		})

		if err != nil {
			return nil, errors.New(consts.InternalServerError)
		}

	}

	//updating cart total price
	cart.Total = subTotal

	_, err = p.cartRepository.UpdateCart(cart)

	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	cart, err = p.cartRepository.FetchCart(currentUserId)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	return cart, nil
}
