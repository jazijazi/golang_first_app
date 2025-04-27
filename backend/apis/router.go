package apis

import (
	// "net/http"
	authApi "httpproj1/auth"
	shopApi "httpproj1/shop/apis"

	"github.com/labstack/echo/v4"
)

func GetRouter() *echo.Echo {
	e := echo.New()

	productRouter := e.Group("/products/")
	productRouter.GET("/find", shopApi.GetProduct)
	productRouter.GET("", shopApi.ListProduct)
	productRouter.POST("", shopApi.CreateProduct)

	brandRouter := e.Group("/brands/")
	brandRouter.GET("", shopApi.ListBrand)
	brandRouter.POST("", shopApi.CreateBrand)

	UserRouter := e.Group("/users/")
	UserRouter.GET("", authApi.ListUser)
	UserRouter.POST("", authApi.CreateUser)
	UserRouter.POST("login/", authApi.Login)
	UserRouter.POST("verify/", authApi.Verify)

	return e
}
