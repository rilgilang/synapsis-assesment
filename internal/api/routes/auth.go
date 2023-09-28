package routes

import (
	"github.com/gofiber/fiber/v2"
	"synapsis-challenge/internal/api/handlers"
	"synapsis-challenge/internal/service"
)

// LoginRouter is the Router for GoFiber App
func LoginRouter(app fiber.Router, service service.AuthService) {
	app.Post("/login", handlers.Login(service))
	app.Post("/register", handlers.Register(service))
}
