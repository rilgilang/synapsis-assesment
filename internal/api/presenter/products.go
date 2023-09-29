package presenter

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/entities"
)

type Products struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
}

func ProductsSuccessResponse(src *[]entities.Product) *fiber.Map {
	var data []Products

	for _, v := range *src {
		p := Products{
			ID:          v.ID,
			ProductName: v.ProductName,
			Category:    v.ProductCategory.CategoryName,
			Price:       v.Price,
		}

		data = append(data, p)
	}

	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func ProductErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
