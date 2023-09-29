package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"synapsis-challenge/internal/api/presenter"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/helper"
	"synapsis-challenge/internal/service"
)

func GetTransactions(service service.TransactionsService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		transactions, err := service.FetchAllTransactions(helper.InterfaceToString(c.Locals(consts.UserId)))
		//only internal server error
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.TransactionsSuccessResponse(transactions))
	}
}
