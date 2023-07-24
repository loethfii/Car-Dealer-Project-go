package models

import "gorm.io/gorm"

type SalesPeople struct {
	gorm.Model
	Id            int            `gorm:"column:id; type: int; autoIncrement" json:"id"`
	NamaSales     string         `gorm:"column:nama_sales; type: varchar(100); not null; " json:"nama_sales"`
	Nip           string         `gorm:"column:nip; type: varchar(50); not null; " json:"nip"`
	NomerTelpon   string         `gorm:"column:nomer_telpon; type: varchar(50); not null; " json:"nomer_telpon"`
	Bagian        string         `gorm:"column:bagian; type: varchar(50); not null; " json:"bagian"`
	PurchaseForms []PurchaseForm `json:"purchase_form"`
	DivisionID    uint           `gorm:"column:division_id; type:uint;" json:"division_id"`
}

type SalesPeopleRequest struct {
	NamaSales   string `json:"nama_sales" binding:"required"`
	Nip         string `json:"nip" binding:"required"`
	NomerTelpon string `json:"nomer_telpon" binding:"required"`
	Bagian      string `json:"bagian" binding:"required"`
	DivisionID  uint   `json:"division_id"`
}

type SalesPeopleResponse struct {
	Id            uint                   `json:"id"`
	NamaSales     string                 `json:"nama_sales"`
	Nip           string                 `json:"nip"`
	NomerTelpon   string                 `json:"nomer_telpon"`
	Bagian        string                 `json:"bagian"`
	PurchaseForms []PurchaseFormResponse `json:"purchase_forms"`
	DivisionID    uint                   `json:"division_id"`
}

type SalesPeopleResponseToPurchaseForm struct {
	Id          uint   `json:"id"`
	NamaSales   string `json:"nama_sales"`
	Nip         string `json:"nip"`
	NomerTelpon string `json:"nomer_telpon"`
	Bagian      string `json:"bagian"`
}
