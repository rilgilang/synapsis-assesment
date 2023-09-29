package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"synapsis-challenge/internal/api/presenter"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/helper"
	"synapsis-challenge/internal/service"
)

func Checkout(service service.CheckoutService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		//calling the service or the bussines logic
		transactions, err := service.CheckoutCart(helper.InterfaceToString(c.Locals(consts.UserId)))
		//only internal server error
		if err != nil {
			if err.Error() == consts.EmptyCart {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(presenter.CartErrorResponse(err))
			}
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.CartErrorResponse(err))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.TransactionsSuccessResponse(transactions))
	}
}
