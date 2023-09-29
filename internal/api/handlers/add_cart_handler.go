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
		//create new struct for our req body
		var params request_model.AddToCart

		//parsing to our struct that we already made before
		err := c.BodyParser(&params)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.CartErrorResponse(err))
		}

		//we need to check if the product id is uuidv4 by using this regex
		compile, _ := regexp.Compile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

		//validation using ozoo
		err = validation.ValidateStruct(&params,
			validation.Field(&params.ProductId, validation.Required, validation.Length(36, 36), validation.Match(compile)),
			validation.Field(&params.Total, validation.Required, validation.Min(0)),
		)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.CartErrorResponse(err))
		}

		//calling the service or the bussines logic
		cart, err := service.AddToCart(helper.InterfaceToString(c.Locals(consts.UserId)), params)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.CartErrorResponse(err))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.CartSuccessResponse(cart))
	}
}
