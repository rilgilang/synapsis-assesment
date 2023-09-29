package presenter

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/entities"
)

type Cart struct {
	ID       string        `json:"id"`
	UserId   string        `json:"user_id"`
	SubTotal int           `json:"sub_total"`
	Products []ProductData `json:"products"`
}

type ProductData struct {
	ID        string `json:"product_cart_id"`
	CartID    string `json:"cart_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Total     int    `json:"total"`
}

func CartSuccessResponse(src *entities.Cart) *fiber.Map {
	productList := []ProductData{}

	for _, v := range src.CartProduct {

		productList = append(productList, ProductData{
			ID:        v.ID,
			CartID:    v.CartId,
			ProductId: v.ProductId,
			Quantity:  v.Quantity,
			Total:     v.Total,
		})
	}

	data := Cart{
		ID:       src.ID,
		UserId:   src.UserId,
		SubTotal: src.Total,
		Products: productList,
	}

	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func CartErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
