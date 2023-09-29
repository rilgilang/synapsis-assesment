package handlers

import (
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

func AddToCart(service service.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params request_model.AddToCart

		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.CartErrorResponse(err))
		}

		compile, _ := regexp.Compile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

		err = validation.ValidateStruct(&params,
			validation.Field(&params.ProductId, validation.Required, validation.Length(36, 36), validation.Match(compile)),
			validation.Field(&params.Total, validation.Required, validation.Min(0)),
		)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.CartErrorResponse(err))
		}

		cart, err := service.AddToCart(helper.InterfaceToString(c.Locals(consts.UserId)), params)
		if err != nil {
			return err
		}

		c.Status(200)
		return c.JSON(presenter.CartSuccessResponse(cart))
	}
}
