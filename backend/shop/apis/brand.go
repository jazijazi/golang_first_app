package apis

import (
	"httpproj1/initializers"
	"httpproj1/shop"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListBrand(c echo.Context) error {
	db := initializers.DB

	var brands []shop.Brand
	db.Find(&brands)
	return c.JSON(http.StatusOK, brands)
}

func CreateBrand(c echo.Context) error {
	db := initializers.DB
	var brand shop.Brand

	if err := c.Bind(&brand); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if result := db.Create(&brand); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, brand)
}
