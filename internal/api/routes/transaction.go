package routes

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/api/handlers"
	"synapsis-challenge/internal/middlewares/jwt"
	"synapsis-challenge/internal/service"
)

func TransactionsRouter(app fiber.Router, middleware jwt.AuthMiddleware, transactionService service.TransactionsService) {
	app.Get("/transactions", middleware.ValidateToken(), handlers.GetTransactions(transactionService))
}
