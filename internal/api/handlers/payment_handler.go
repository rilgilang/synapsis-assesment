package handlers

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"regexp"
	"synapsis-challenge/internal/api/presenter"
	"synapsis-challenge/internal/api/request_model"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/helper"
	"synapsis-challenge/internal/service"
)

func Payment(service service.TransactionsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params request_model.PaymentTransaction

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TransactionsErrorResponse(errors.New(consts.BadRequest)))
		}

		compile, _ := regexp.Compile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

		err = validation.ValidateStruct(&params,
			validation.Field(&params.TransactionId, validation.Required, validation.Length(36, 36), validation.Match(compile)),
			validation.Field(&params.Amount, validation.Required, validation.Min(0)),
		)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.TransactionsErrorResponse(err))
		}

		transaction, err := service.PaymentTransaction(helper.InterfaceToString(c.Locals(consts.UserId)), params)
		if err != nil {
			if err.Error() == consts.InsufficientAmount {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(presenter.TransactionsErrorResponse(err))
			}

			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.TransactionsErrorResponse(err))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.TransactionPaymentSuccessResponse(transaction))
	}
}
