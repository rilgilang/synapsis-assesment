package presenter

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/entities"
)

type Transactions struct {
	ID       string               `json:"id"`
	SubTotal int                  `json:"sub_total"`
	Status   string               `json:"status"`
	Products []TransactionProduct `json:"products"`
}

type TransactionProduct struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Total       int    `json:"total"`
}

func TransactionPaymentSuccessResponse(src *entities.Transaction) *fiber.Map {
	//TODO find better way to present transactions
	data := Transactions{
		ID:       src.ID,
		SubTotal: src.Total,
		Status:   src.Status,
		Products: nil,
	}

	var products []TransactionProduct
	for _, product := range src.Order {
		pr := TransactionProduct{
			ID:          product.ProductId,
			ProductName: product.ProductName,
			Quantity:    product.Quantity,
			Total:       product.Total,
		}
		products = append(products, pr)
	}

	data.Products = products

	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func TransactionsSuccessResponse(src *[]entities.Transaction) *fiber.Map {
	data := []Transactions{}
	//TODO find better way to present transactions
	for _, transaction := range *src {
		tr := Transactions{
			ID:       transaction.ID,
			SubTotal: transaction.Total,
			Status:   transaction.Status,
			Products: nil,
		}

		products := []TransactionProduct{}
		for _, product := range transaction.Order {
			pr := TransactionProduct{
				ID:          product.ProductId,
				ProductName: product.ProductName,
				Quantity:    product.Quantity,
				Total:       product.Total,
			}
			products = append(products, pr)
		}

		tr.Products = products

		data = append(data, tr)
	}

	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func TransactionsErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
