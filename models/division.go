package models

import "gorm.io/gorm"

type Division struct {
	gorm.Model
	NameDivision string        `gorm:"column:name_division; type: varchar(100);" json:"name_division"`
	Code         string        `gorm:"column:code; type: varchar(100);" json:"code"`
	DepartmentID uint          `gorm:"column:department_id; type: uint;" json:"department_id"`
	SalesPeoples []SalesPeople `json:"sales_people"`
}

type DivisionRequest struct {
	NameDivision string `json:"name_division"`
	Code         string `json:"code"`
	DepartmentID uint   `json:"department_id"`
}

type DivisionResponse struct {
	ID           uint   `json:"id"`
	NameDivision string `json:"name_division"`
	Code         string `json:"code"`
	DepartmentID uint   `json:"department_id"`
}
