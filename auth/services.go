package auth

import "httpproj1/initializers"

func createUserToDb(userRequest UserRequest) User {
	db := initializers.DB

	user := User{
		Name:     userRequest.Name,
		Password: userRequest.Password,
		Role:     userRequest.Role,
	}

	db.Create(&user)

	return user
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
