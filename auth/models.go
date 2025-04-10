package auth

import (
	// "gorm.io/driver/postgres"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name     string    ``
	Role     string    ``
	Password string    `gorm:"type:varchar(50)"`
}
