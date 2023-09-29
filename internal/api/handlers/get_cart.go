package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"synapsis-challenge/internal/api/presenter"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/helper"
	"synapsis-challenge/internal/service"
)

func GetCart(service service.CartService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		//bussines logic
		cart, err := service.FetchCart(helper.InterfaceToString(c.Locals(consts.UserId)))
		//only internal server error
		if err != nil && err.Error() == consts.InternalServerError {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.CartErrorResponse(err))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.CartSuccessResponse(cart))
	}

}
