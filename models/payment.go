package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Id             int    `gorm:"column:id; type: int; autoIncrement" json:"id"`
	BuktiTransfer  string `gorm:"column:bukti_transfer; type: varbinary(100); " json:"bukti_transfer"`
	IsConfirm      bool   `gorm:"column:is_confirm; type: bool;default: false;" json:"is_confirm"`
	PurchaseFormId int    `gorm:"column:purchase_form_id;type:int;not null;unique" json:"purchase_form_id"`
}

type PaymentRequest struct {
	BuktiTransfer  string `json:"bukti_transfer" form:"bukti_transfer"`
	IsConfirm      bool   `json:"is_confirm" form:"is_confirm"`
	PurchaseFormId int    `json:"purchase_form_id" form:"purchase_form_id"`
}

type PaymentResponse struct {
	Id             int
	BuktiTransfer  string `json:"bukti_transfer" form:"bukti_transfer"`
	IsConfirm      bool   `json:"is_confirm" form:"bukti_transfer"`
	PurchaseFormId int    `json:"purchase_form_id" form:"purchase_form_id"`
}
