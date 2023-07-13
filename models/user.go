package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Username string `gorm:"column:username;type:varchar(50);unique;not null" json:"Username"`
	Email    string `gorm:"column:email;type:varchar(50);unique;not null" json:"email"`
	Password string `gorm:"column:password;type:varchar(200);not null" json:"password"`
	Role     int    `gorm:"column:role;type:int;not null; DEFAULT : 3" json:"role"`
}

type TokenRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Username string `json:"Username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}
