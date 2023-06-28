package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `gorm:"column:id;type: int; autoIncrement" json:"Id"`
	Username string `gorm:"column:username;type:varchar(50);unique;not null" json:"Username"`
	Password string `gorm:"column:password;type:varchar(50);not null" json:"password"`
	Role     int    `gorm:"column:role;type:varchar(10);not null; DEFAULT : 3" json:"role"`
}
