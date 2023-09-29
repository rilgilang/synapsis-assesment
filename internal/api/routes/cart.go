package routes

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/api/handlers"
	"synapsis-challenge/internal/middlewares/jwt"
	"synapsis-challenge/internal/service"
)

func CartRouter(app fiber.Router, middleware jwt.AuthMiddleware, cartService service.CartService) {
	app.Post("/cart/add", middleware.ValidateToken(), handlers.AddToCart(cartService))
	app.Get("/cart", middleware.ValidateToken(), handlers.GetCart(cartService))
}
