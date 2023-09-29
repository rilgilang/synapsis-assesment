package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"synapsis-challenge/internal/api/presenter"
	"synapsis-challenge/internal/api/request_model"
	"synapsis-challenge/internal/consts"
	"synapsis-challenge/internal/service"
)

func GetProducts(service service.ProductService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params request_model.GetProducts

		params.Category = c.Query("category")

		err := validation.ValidateStruct(&params,
			validation.Field(&params.Category, validation.Length(0, 20)),
		)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		products, err := service.FetchAllProduct(params)
		//only internal server error
		if err != nil && err.Error() == consts.InternalServerError {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ProductErrorResponse(err))
		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.ProductsSuccessResponse(products))
	}
}
