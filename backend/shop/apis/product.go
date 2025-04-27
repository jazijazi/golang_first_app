package apis

import (
	"httpproj1/shop"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

func ListProduct(c echo.Context) error {
	products, err := shop.ListProductService(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	title := c.QueryParam("title")
	if title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Title query parameter is required",
		})
	}

	product, err := shop.GetProductService(title)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	var product shop.ProductRequest

	// Bind request body to product struct
	if err := c.Bind(&product); err != nil {
		//Use echo.Map instead of map[string]string for JSON responses (it's cleaner with Echo).
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	// Validate the product fields
	if err := validation.ValidateStruct(&product,
		validation.Field(&product.Title, validation.Required, validation.Length(2, 50)),
		validation.Field(&product.Price, validation.Required),
	); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Validation failed: " + err.Error(),
		})
	}

	// Call service to create product
	if err := shop.CreateProductService(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create product: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
	})
}
