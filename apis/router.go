package apis

import (
	// "net/http"
	"github.com/labstack/echo/v4"
)

func GetRouter() *echo.Echo {
	e := echo.New()
	brandRouter := e.Group("/brands/")
	brandRouter.GET("", listBrand)
	brandRouter.POST("", createBrand)
	return e
}
