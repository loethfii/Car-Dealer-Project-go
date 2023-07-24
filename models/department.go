package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	NameDepartment string     `gorm:"column:name_department; type: varchar(100);" json:"name_department"`
	Code           string     `json:"code" gorm:"column:code; type: varchar(100);"`
	Divisions      []Division `json:"division"`
}

type DepartmentRequest struct {
	NameDepartment string `json:"name_department"`
	Code           string `json:"code"`
}

type DepartementResponse struct {
	ID             uint               `json:"id"`
	NameDepartment string             `json:"name_department"`
	Code           string             `json:"code"`
	Divisions      []DivisionResponse `json:"division"`
}
