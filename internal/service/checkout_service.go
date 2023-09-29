package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/entities"
	"synapsis-challenge/internal/repositories"
	"time"
)

type CheckoutService interface {
	CheckoutCart(curentUserId string) (*[]entities.Transaction, error)
}

type checkoutService struct {
	transactionRepository repositories.TransactionsRepository
	cartRepository        repositories.CartRepository
	orderRepository       repositories.OrderRepository
	productCartRepository repositories.CartProductRepository
}

func NewCheckoutService(
	transactionRepository repositories.TransactionsRepository,
	cartRepository repositories.CartRepository,
	orderRepository repositories.OrderRepository,
	productCartRepository repositories.CartProductRepository,
) CheckoutService {
	return &checkoutService{
		transactionRepository: transactionRepository,
		cartRepository:        cartRepository,
		orderRepository:       orderRepository,
		productCartRepository: productCartRepository,
	}
}

func (p checkoutService) CheckoutCart(currentUserId string) (*[]entities.Transaction, error) {

	//get sub total price from cart
	cart, err := p.cartRepository.FetchCart(currentUserId)
	if err != nil {
		return nil, err
	}

	//check if cart has products in it
	if len(cart.CartProduct) == 0 {
		return nil, errors.New(consts.EmptyCart)
	}

	//create new transaction
	transactions, err := p.transactionRepository.NewTransactions(&entities.Transaction{
		ID:        uuid.New().String(),
		UserId:    currentUserId,
		Total:     cart.Total,
		Status:    "waiting",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	//then insert products to order
	for _, product := range cart.CartProduct {
		err = p.orderRepository.InsertOrder(&entities.Order{
			ID:            uuid.New().String(),
			TransactionId: transactions.ID,
			ProductId:     product.ProductId,
			ProductName:   product.ProductName,
			Quantity:      product.Quantity,
			Total:         product.Total,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		})

		if err != nil {
			return nil, err
		}

	}

	cart.Total = 0
	//cleaning cart
	//update cart total
	_, err = p.cartRepository.UpdateCart(cart)
	if err != nil {
		return nil, err
	}

	//delete all product from cart
	err = p.productCartRepository.DeleteAllProductFromCart(cart.ID, currentUserId)
	if err != nil {
		return nil, err
	}

	transactionList, err := p.transactionRepository.FetchTransactions(currentUserId)
	if err != nil {
		return nil, err
	}

	return transactionList, nil
}
