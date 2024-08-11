package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	FullName string `json: "fullname"`
	Email    string `json: "email" validate:"required"`
	Password string `json: "password" validate:"required"`
	Phone    string `json: "phone" validate: "phone"`
}
