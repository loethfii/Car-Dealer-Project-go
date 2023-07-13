package models

import "gorm.io/gorm"

type PurchaseForm struct {
	gorm.Model
	Id                 int     `gorm:"column:id; type: int; autoIncrement" json:"id"`
	NamaLengkapPembeli string  `gorm:"column:nama_lengkap_pembeli; type: varchar(100); " json:"nama_lengkap_pembeli"`
	NomerKTP           string  `gorm:"column:nomer_ktp; type: varchar(100); " json:"nomer_ktp"`
	AlamatRumah        string  `gorm:"column:alamat_rumah; type: varchar(100); " json:"alamat_rumah"`
	NomerDebit         string  `gorm:"column:nomer_debit; type: varchar(100); " json:"nomer_debit"`
	CarId              int     `gorm:"column:car_id; type: int; " json:"car_id"`
	HarusInden         bool    `gorm:"column:harus_inden; type: bool; " json:"harus_inden"`
	LamaInden          string  `gorm:"column:lama_inden; type: varchar(100); " json:"lama_inden"`
	CustomPlat         string  `gorm:"column:custom_plat; type: varchar(100); " json:"custom_plat"`
	TambahanKit        string  `gorm:"column:tambahan_kit; type: varchar(100); " json:"tambahan_kit"`
	SalesPeopleId      int     `gorm:"column:sales_people_id;type: int; " json:"sales_people_id"`
	Payment            Payment `gorm:"foreignKey:PurchaseFormId;"`
}

type PurchaseFormRequest struct {
	NamaLengkapPembeli string `json:"nama_lengkap_pembeli" binding:"required"`
	NomerKTP           string `json:"nomer_ktp" binding:"required"`
	AlamatRumah        string `json:"alamat_rumah" binding:"required"`
	NomerDebit         string `json:"nomer_debit" binding:"required"`
	CarId              int    `json:"car_id" binding:"required"`
	HarusInden         bool   `json:"harus_inden"`
	LamaInden          string `json:"lama_inden" `
	CustomPlat         string `json:"custom_plat"`
	TambahanKit        string `json:"tambahan_kit"`
	SalesPeopleId      int    `json:"sales_people_id" binding:"required"`
}

type PurchaseFormResponse struct {
	Id                 int    `json:"id"`
	NamaLengkapPembeli string `json:"nama_lengkap_pembeli"`
	NomerKTP           string `json:"nomer_ktp"`
	AlamatRumah        string `json:"alamat_rumah"`
	NomerDebit         string `json:"nomer_debit"`
	CarId              int    `json:"car_id"`
	HarusInden         bool   `json:"harus_inden"`
	LamaInden          string `json:"lama_inden"`
	CustomPlat         string `json:"custom_plat"`
	TambahanKit        string `json:"tambahan_kit"`
	SalesPeopleId      int    `json:"sales_people_id"`
}

type PurchaseFormInnerJoinResponse struct {
	Id                 int    `json:"id"`
	NamaLengkapPembeli string `json:"nama_lengkap_pembeli"`
	NomerKTP           string `json:"nomer_ktp"`
	AlamatRumah        string `json:"alamat_rumah"`
	NomerDebit         string `json:"nomer_debit"`
	CarId              int    `json:"car_id"`
	CarDetail          CarResponseToPurchaseForm
	HarusInden         bool   `json:"harus_inden"`
	LamaInden          string `json:"lama_inden"`
	CustomPlat         string `json:"custom_plat"`
	TambahanKit        string `json:"tambahan_kit"`
	SalesPeopleId      int    `json:"sales_people_id"`
	SalesPeopleDetail  SalesPeopleResponseToPurchaseForm
	PaymentID          int
	BuktiTransfer      string `json:"bukti_transfer"`
	IsConfirm          bool   `json:"is_confirm"`
}
