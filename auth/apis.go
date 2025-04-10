package auth

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

func ListUser(c echo.Context) error {
	users := getUsers()
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	// db := initializers.DB
	var user UserRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	valiadtion_err := validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&user.Password, validation.Required, validation.Length(2, 5)),
		validation.Field(&user.Role),
	)
	if valiadtion_err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": valiadtion_err.Error()})
	}

	// if result := db.Create(&user); result.Error != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	// }
	user_obj := createUserToDb(user)

	return c.JSON(http.StatusOK, user_obj)
}
