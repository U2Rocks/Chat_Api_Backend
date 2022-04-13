package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
}
