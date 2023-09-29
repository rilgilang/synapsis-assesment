package routes

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/api/handlers"
	"synapsis-challenge/internal/middlewares/jwt"
	"synapsis-challenge/internal/service"
)

func CheckoutRouter(app fiber.Router, middleware jwt.AuthMiddleware, checkoutService service.CheckoutService) {
	app.Post("/checkout", middleware.ValidateToken(), handlers.Checkout(checkoutService))
}
