package repository

import (
	"errors"
	"gorm.io/gorm"
	"tugas_akhir_course_net/models"
)

type PurchaseFormRepository interface {
	FindAll() ([]models.PurchaseForm, error)
	FindById(id int) (models.PurchaseForm, error)
	Create(purchaseForm models.PurchaseForm, id int, isInden bool) (models.PurchaseForm, error)
	Update(id int, updatePurchaseForm models.PurchaseFormRequest) (models.PurchaseFormRequest, error)
	Delete(id int) (models.PurchaseForm, error)
}

type purchaseFormRepository struct {
	db *gorm.DB
}

func NewPurchaseFormRepository(db *gorm.DB) *purchaseFormRepository {
	return &purchaseFormRepository{db}
}

func (r *purchaseFormRepository) FindAll() ([]models.PurchaseForm, error) {
	var purchaseForm []models.PurchaseForm

	err := r.db.Order("id desc").Find(&purchaseForm).Error

	return purchaseForm, err
}

func (r *purchaseFormRepository) FindById(id int) (models.PurchaseForm, error) {
	var purchaseForm models.PurchaseForm

	err := r.db.First(&purchaseForm, id).Error

	return purchaseForm, err
}

func (r *purchaseFormRepository) Create(purchaseForm models.PurchaseForm, id int, isInden bool) (models.PurchaseForm, error) {

	var err error
	var cars models.Car

	err = r.db.Create(&purchaseForm).Error

	result := r.db.First(&cars, id)

	if cars.Qty == 0 && isInden == false {
		err = errors.New("Mobil sedang Kosong tidak dapat dibeli instan harus inden!")
	}

	if cars.Qty >= 1 && isInden == true {
		err = errors.New("Mobil tersedia di bagasi mohon inden di false kan! ")
	}

	if result.RowsAffected == 0 {
		err = errors.New("Data Mobil Tidak Tersedia")
	}

	return purchaseForm, err
}

func (r *purchaseFormRepository) Update(id int, updatePurchaseForm models.PurchaseFormRequest) (models.PurchaseFormRequest, error) {

	err := r.db.Model(&models.PurchaseForm{}).Where("id = ?", id).Updates(&updatePurchaseForm).Error

	return updatePurchaseForm, err
}

func (r *purchaseFormRepository) Delete(id int) (models.PurchaseForm, error) {

	var purchaseForm models.PurchaseForm

	r.db.First(&purchaseForm, id)

	err := r.db.Delete(&purchaseForm).Error

	return purchaseForm, err
}
