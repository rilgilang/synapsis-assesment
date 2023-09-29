package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"net/http"
	"synapsis-challenge/internal/api/presenter"
	"synapsis-challenge/internal/api/request_model"
	"synapsis-challenge/internal/consts"
	entities2 "synapsis-challenge/internal/entities"
	"synapsis-challenge/internal/service"
)

func Login(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request_model.Login
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.AuthErrorResponse(err))
		}
		if requestBody.Username == "" || requestBody.Password == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.AuthErrorResponse(errors.New(
				"Please specify username and password")))
		}

		userData := entities2.User{Username: requestBody.Username, Password: requestBody.Password}
		user, token, err := service.Login(&userData)
		//only internal server error
		if err != nil && err.Error() == consts.InternalServerError {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		//can be unauthorized or something else
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.AuthSuccesResponse(user, *token))
	}
}

func Register(service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody request_model.Register
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.AuthErrorResponse(err))
		}
		if requestBody.Username == "" || requestBody.Password == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.AuthErrorResponse(errors.New(
				"Please specify username and password")))
		}

		userData := entities2.User{Username: requestBody.Username, Password: requestBody.Password}
		user, token, err := service.Register(&userData)
		//only internal server error
		if err != nil && err.Error() == consts.InternalServerError {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		//can be unauthorized or something else
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.AuthErrorResponse(err))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.AuthSuccesResponse(user, *token))
	}
}
