package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Id            int          `gorm:"column:id; type: int; autoIncrement" json:"id"`
	FormPembelian PurchaseForm `gorm:"column:form_pembelian; type: int; " json:"form_pembelian"`
	BuktiTransfer string       `gorm:"column:bukti_transfer; type: varchar(100); " json:"bukti_transfer"`
	IsConfirm     bool         `gorm:"column:is_confirm; type: bool; " json:"is_confirm"`
}
