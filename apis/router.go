package apis

import (
	// "net/http"
	authApi "httpproj1/auth"
	shopApi "httpproj1/shop/apis"

	"github.com/labstack/echo/v4"
)

func GetRouter() *echo.Echo {
	e := echo.New()
	brandRouter := e.Group("/brands/")
	brandRouter.GET("", shopApi.ListBrand)
	brandRouter.POST("", shopApi.CreateBrand)

	UserRouter := e.Group("/users/")
	UserRouter.GET("", authApi.ListUser)
	UserRouter.POST("", authApi.CreateUser)

	return e
}
