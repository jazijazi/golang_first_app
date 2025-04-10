package auth

import (
	// "gorm.io/driver/postgres"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name     string    ``
	Role     string    ``
	Password string    `gorm:"type:varchar(200)"`
}

func (user *User) hashPassword() error {
	if len(user.Password) < 5 {
		return errors.New("password must more than 5 letter")
	}
	passwordHash, errorHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errorHash != nil {
		return errorHash
	}
	user.Password = string(passwordHash)
	return nil
}

func (user *User) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(user.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

type MyCustomClaims struct {
	jwt.RegisteredClaims
	Name string `json:"name"`
	Role string `json:"role"`
}
