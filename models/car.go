package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Id            int            `gorm:"column:id;type: int; autoIncrement" json:"id"`
	NamaMobil     string         `gorm:"column:nama_mobil;type: varchar(100); " json:"nama_mobil"`
	TipeMobil     string         `gorm:"column:tipe_mobil;type: varchar(100); " json:"tipe_mobil"`
	JenisMobil    string         `gorm:"column:jenis_mobil;type: varchar(100); " json:"jenis_mobil"`
	BahanBakar    string         `gorm:"column:bahan_bakar;type: varchar(100); " json:"bahan_bakar"`
	Isi_Silinder  string         `gorm:"column:isi_Silinder;type: varchar(100); " json:"isi_Silinder"`
	Warna         string         `gorm:"column:warna;type: varchar(100); " json:"warna"`
	Transmisi     string         `gorm:"column:transmisi;type: varchar(100); " json:"transmisi"`
	Harga         uint32         `gorm:"column:harga;type: int; " json:"harga"`
	Qty           uint32         `gorm:"column:qty;type: int; " json:"qty"`
	PurchaseForms []PurchaseForm `gorm:"foreignKey:CarId"`
}

type CarResponse struct {
	Id            int    `json:"id"`
	NamaMobil     string `json:"nama_mobil"`
	TipeMobil     string `json:"tipe_mobil"`
	JenisMobil    string `json:"jenis_mobil"`
	BahanBakar    string `json:"bahan_bakar"`
	Isi_Silinder  string `json:"isi_Silinder"`
	Warna         string `json:"warna"`
	Transmisi     string `json:"transmisi"`
	Harga         uint32 `json:"harga"`
	Qty           uint32 `json:"qty"`
	PurchaseForms []PurchaseFormResponse
}

type CarRequest struct {
	NamaMobil    string `json:"nama_mobil"`
	TipeMobil    string `json:"tipe_mobil"`
	JenisMobil   string `json:"jenis_mobil"`
	BahanBakar   string `json:"bahan_bakar"`
	Isi_Silinder string `json:"isi_Silinder"`
	Warna        string `json:"warna"`
	Transmisi    string `json:"transmisi"`
	Harga        uint32 `json:"harga"`
	Qty          uint32 `json:"qty"`
}
