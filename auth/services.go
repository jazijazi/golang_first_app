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
		// myslog.Error("Error in Load Config File!")
		fmt.Println(err.Error())
	}

	if checkResultError := user.checkPassword(request.Password); checkResultError != nil {
		return LoginResponse{}, checkResultError
	} else {
		// return LoginResponse{Token: "oooooooooooo"}, err
		claims := MyCustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			},
			Name: user.Name,
			Role: user.Role,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		accessToken, accessTokenError := token.SignedString([]byte(config.SECRETKEY))
		if accessTokenError != nil {
			return LoginResponse{}, accessTokenError
		}
		return LoginResponse{Token: accessToken}, nil

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
