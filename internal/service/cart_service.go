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
	FetchCart(currentUserId string) (*entities.Cart, error)
	DeleteFromCart(currentUserId string, param request_model.DeleteFromCart) error
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

func (p *cartService) AddToCart(currentUserId string, param request_model.AddToCart) (*entities.Cart, error) {
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

	//this is used for checking if product already in cart or not
	checkProduct := map[string]string{}

	//initialize product not already in cart
	checkProduct["alreadyInCart"] = "false"

	//set the quantity from the param
	newCartProductQuantity := param.Total

	//set default total price
	newCartProductTotal := 0
	subTotal := 0

	//iterate through product that already in cart
	for _, productCart := range *productsInCart {

		//if product already in cart
		if productCart.ProductId == param.ProductId {

			//set true we use this for validation later
			checkProduct["alreadyInCart"] = "true"
			checkProduct["productCartId"] = productCart.ID

			//set the product total price
			newCartProductTotal = newCartProductQuantity * productsData.Price

			//basic formula for sub total
			subTotal += newCartProductTotal
		} else {
			subTotal += newCartProductTotal
			subTotal += productCart.Total
		}
	}

	//if product already in cart then we dont need to insert just update the product based on the request body
	if checkProduct["alreadyInCart"] == "true" {
		_, err = p.productCartRepository.UpdateProductInCart(cart.ID, checkProduct["productCartId"], newCartProductQuantity, newCartProductTotal)

		if err != nil {
			return nil, errors.New(consts.InternalServerError)
		}

	} else {

		//if product not in cart then insert

		//inserting product to cart
		err = p.productCartRepository.ProductInsertToCart(&entities.CartProduct{
			ID:          uuid.New().String(),
			UserId:      currentUserId,
			CartId:      cart.ID,
			ProductId:   param.ProductId,
			ProductName: productsData.ProductName,
			Quantity:    param.Total,
			Total:       param.Total * productsData.Price,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		})

		if err != nil {
			return nil, errors.New(consts.InternalServerError)
		}

		subTotal += param.Total * productsData.Price
	}

	//updating cart total price
	cart.Total = subTotal

	_, err = p.cartRepository.UpdateCart(cart)

	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	//fetch latest cart data for the response
	cart, err = p.cartRepository.FetchCart(currentUserId)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	return cart, nil
}

func (p *cartService) FetchCart(currentUserId string) (*entities.Cart, error) {
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

	}

	cart, err = p.cartRepository.FetchCart(currentUserId)
	if err != nil {
		return nil, errors.New(consts.InternalServerError)
	}

	return cart, nil
}

func (p *cartService) DeleteFromCart(currentUserId string, param request_model.DeleteFromCart) error {
	//find user cart
	cart, err := p.cartRepository.CheckCart(currentUserId)

	if err != nil {
		return errors.New(consts.InternalServerError)
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
			return errors.New(consts.InternalServerError)
		}

		cart, err = p.cartRepository.CheckCart(currentUserId)

		if err != nil {
			return errors.New(consts.InternalServerError)
		}

		return nil
	}

	//get the product from cart
	//we're gonna grab the total price of this product (product tha already in cart)
	//then total in cart devide with current total price of this product (product that already in cart)
	productInCart, err := p.productCartRepository.FindOneProductInCart(currentUserId, param.Id, cart.ID)
	if err != nil {
		return errors.New(consts.InternalServerError)
	}

	if productInCart == nil {
		return errors.New(consts.NotFound)
	}

	//delete product from cart
	err = p.productCartRepository.DeleteProductFromCart(cart.ID, param.Id, currentUserId)

	//updating cart total price
	cart.Total -= productInCart.Total

	_, err = p.cartRepository.UpdateCart(cart)

	if err != nil {
		return errors.New(consts.InternalServerError)
	}

	return nil
}
