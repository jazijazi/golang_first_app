package auth

import (
	"fmt"
	"httpproj1/initializers"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createUserToDb(userRequest UserRequest) (User, error) {
	db := initializers.DB

	user := User{
		Name:     userRequest.Name,
		Password: userRequest.Password,
		Role:     userRequest.Role,
	}

	hashError := user.hashPassword()
	if hashError != nil {
		return User{}, hashError

	}

	result := db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil

	// if result := db.Create(&user); result.Error != nil {
	// return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	// }
}

func getUsers() []User {
	db := initializers.DB
	var users []User
	db.Find(&users)
	return users
}

func getUserByName(name string) (User, error) {
	db := initializers.DB
	var user User
	if result := db.Where("name = ?", name).First(&user); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func login(request LoginRequest) (LoginResponse, error) {
	user, dberr := getUserByName(request.Name)

	if dberr != nil {
		return LoginResponse{}, dberr
	}

	config, err := initializers.LoadConfig(".")
	if err != nil {
		fmt.Println(err.Error())
	}

	if checkResultError := user.checkPassword(request.Password); checkResultError != nil {
		return LoginResponse{}, checkResultError
	} else {
		// Create Claims for Access Token
		claims := MyCustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			},
			Name: user.Name,
			Role: user.Role,
		}
		// Create Access Token
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		accessTokenStr, accessTokenError := accessToken.SignedString([]byte(config.SECRETKEY))
		if accessTokenError != nil {
			return LoginResponse{}, accessTokenError
		}

		// Create Refresh Token with a longer expiration
		refreshClaims := MyCustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // Set expiration time (1 week)
			},
			Name: user.Name,
			Role: user.Role,
		}
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
		refreshTokenStr, refreshTokenError := refreshToken.SignedString([]byte(config.SECRETKEY))
		if refreshTokenError != nil {
			return LoginResponse{}, refreshTokenError
		}

		// Return Access Token and Refresh Token
		return LoginResponse{AccessToken: accessTokenStr, RefreshToken: refreshTokenStr}, nil
	}
}

func verify(request VerifyRequest) (VerifyResponse, error) {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		// myslog.Error("Error in Load Config File!")
		fmt.Println(err.Error())
	}

	token, err := jwt.ParseWithClaims(request.Token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRETKEY), nil
	})

	if clm, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Println(clm)
		return VerifyResponse{IsValid: true}, err
	}
	return VerifyResponse{IsValid: false}, err

}

func refresh(refreshtoken string) (string, error) {
	// Parse and validate the refresh token
	config, _ := initializers.LoadConfig(".")
	claims := &MyCustomClaims{}

	token, err := jwt.ParseWithClaims(refreshtoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SECRETKEY), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid or expired refresh token")
	}

	userName := claims.Name
	userRole := claims.Role

	fmt.Println("========")
	fmt.Println(claims.Name)
	fmt.Println(claims.Role)
	fmt.Println("========")

	// If the refresh token is valid, create a new access token
	newClaims := MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), // New access token valid for 2 hours
		},
		Name: userName,
		Role: userRole,
	}
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newAccessTokenStr, err := newAccessToken.SignedString([]byte(config.SECRETKEY))
	if err != nil {
		return "", fmt.Errorf("error generating new access token")
	}
	return newAccessTokenStr, nil
}
