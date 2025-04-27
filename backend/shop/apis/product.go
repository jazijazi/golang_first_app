package apis

import (
	"httpproj1/shop"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

func ListProduct(c echo.Context) error {
	products := shop.ListProductService()
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	var product shop.ProductRequest

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	validationErr := validation.ValidateStruct(&product,
		validation.Field(&product.Title, validation.Required, validation.Length(2, 50)),
		validation.Field(&product.Price, validation.Required),
	)

	if validationErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": validationErr.Error()})
	}

	db_err := shop.CreateProductService(&product)

	if db_err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": db_err.Error()})
	}

	return c.JSON(http.StatusOK, true)
}

func GetProduct(c echo.Context) error {
	title := c.QueryParam("title")
	product := shop.GetProductService(title)
	return c.JSON(http.StatusOK, product)
}
