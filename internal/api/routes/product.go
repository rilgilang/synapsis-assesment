package routes

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/api/handlers"
	"synapsis-challenge/internal/service"
)

func ProductRouter(app fiber.Router, productService service.ProductService) {
	app.Get("/products", handlers.GetProducts(productService))
}
