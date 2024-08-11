package models

import "gorm.io/gorm"

type Patients struct {
	gorm.Model
	FullName string `json: "fullname"`
	Email    string `json: "email"`
	Dni      int    `json: "dni"`
	Phone    int    `json: "phone"`
}
