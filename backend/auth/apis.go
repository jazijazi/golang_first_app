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
		validation.Field(&user.Password, validation.Required, validation.Length(5, 25)),
		validation.Field(&user.Role),
	)
	if valiadtion_err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": valiadtion_err.Error()})
	}

	// if result := db.Create(&user); result.Error != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	// }
	user_obj, err := createUserToDb(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user_obj)
}

func Login(c echo.Context) error {
	var request LoginRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	valiadtion_err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&request.Password, validation.Required, validation.Length(5, 25)),
	)
	if valiadtion_err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": valiadtion_err.Error()})
	}

	// Call Login function
	loginResponse, loginErr := login(request)
	if loginErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": loginErr.Error()})
	}

	// Set Refresh Token as HTTP-only cookie
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    loginResponse.RefreshToken,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7, // 7 days
		HttpOnly: true,
		Secure:   false, // Set to true if using HTTPS
		SameSite: http.SameSiteStrictMode,
	})
	// Return Access Token in the response body
	return c.JSON(http.StatusOK, map[string]string{
		"token": loginResponse.AccessToken,
	})
}

func Verify(c echo.Context) error {
	var request VerifyRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	valiadtion_err := validation.ValidateStruct(&request,
		validation.Field(&request.Token, validation.Required),
	)
	if valiadtion_err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": valiadtion_err.Error()})
	}

	response, verifyError := verify(request)
	if verifyError != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": verifyError.Error()})
	}
	return c.JSON(http.StatusBadRequest, response)
}

func Refresh(c echo.Context) error {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken.Value == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Refresh token is missing"})
	}

	newAccessToken, refreshErr := refresh(refreshToken.Value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": refreshErr.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": newAccessToken,
	})
}
