package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tugas_akhir_course_net/models"
)

type PurchaseFormRepository interface {
	FindAll() ([]models.PurchaseForm, error)
	FindById(id int) (models.PurchaseForm, int, int, error)
	Create(purchaseForm models.PurchaseForm, carID int, salesID int, isInden bool) (models.PurchaseForm, error)
	Update(id int, updatePurchaseForm models.PurchaseFormRequest) (models.PurchaseFormRequest, error)
	Delete(id int) (models.PurchaseForm, error)
	FindBySalesPeopleID(id int) ([]models.PurchaseForm, error)
	FindByCarID(id int) ([]models.PurchaseForm, error)
}

type purchaseFormRepository struct {
	db *gorm.DB
}

func NewPurchaseFormRepository(db *gorm.DB) *purchaseFormRepository {
	return &purchaseFormRepository{db}
}

func (r *purchaseFormRepository) FindAll() ([]models.PurchaseForm, error) {
	var purchaseForms []models.PurchaseForm

	err := r.db.Order("id desc").Preload(clause.Associations).Find(&purchaseForms).Error

	return purchaseForms, err
}

func (r *purchaseFormRepository) FindById(id int) (models.PurchaseForm, int, int, error) {
	var purchaseForm models.PurchaseForm

	err := r.db.Preload(clause.Associations).First(&purchaseForm, id).Error

	var carId = purchaseForm.CarId
	var salesPeopleId = purchaseForm.SalesPeopleId

	return purchaseForm, carId, salesPeopleId, err
}

func (r *purchaseFormRepository) Create(purchaseForm models.PurchaseForm, carID int, salesID int, isInden bool) (models.PurchaseForm, error) {

	var err error
	var cars models.Car
	resultCar := r.db.First(&cars, carID)

	var sales models.SalesPeople
	resultSales := r.db.First(&sales, salesID)

	if resultSales.RowsAffected == 0 && resultCar.RowsAffected == 0 {
		err = errors.New("Data sales dan mobil tidak ditemukan")
	} else if resultCar.RowsAffected == 0 {
		err = errors.New("Data mobil tidak ditemukan")
	} else if resultSales.RowsAffected == 0 {
		err = errors.New("Data sales tidak ditemukan")
	} else if cars.Qty == 0 && isInden == false {
		err = errors.New("Mobil sedang Kosong tidak dapat dibeli instan harus inden!")
	} else if cars.Qty >= 1 && isInden == true {
		err = errors.New("Mobil tersedia di bagasi mohon inden di false kan! ")
	} else {
		r.db.Create(&purchaseForm)
	}

	return purchaseForm, err
}

func (r *purchaseFormRepository) Update(id int, updatePurchaseForm models.PurchaseFormRequest) (models.PurchaseFormRequest, error) {
	var err error
	result := r.db.Model(&models.PurchaseForm{}).Where("id = ?", id).Updates(&updatePurchaseForm)
	if result.RowsAffected == 0 {
		err = errors.New("data tidak ditemukan")
	}

	return updatePurchaseForm, err
}

func (r *purchaseFormRepository) Delete(id int) (models.PurchaseForm, error) {

	var purchaseForm models.PurchaseForm

	r.db.First(&purchaseForm, id)

	err := r.db.Delete(&purchaseForm).Error

	return purchaseForm, err
}

func (r *purchaseFormRepository) FindBySalesPeopleID(id int) ([]models.PurchaseForm, error) {
	var purchaseForm []models.PurchaseForm
	var err error

	result := r.db.Order("id desc").Preload(clause.Associations).Find(&purchaseForm, "sales_people_id = ?", id)
	if result.RowsAffected == 0 {
		err = errors.New("data tidak ditemukan")
	}

	return purchaseForm, err
}

func (r *purchaseFormRepository) FindByCarID(id int) ([]models.PurchaseForm, error) {
	var purchaseForm []models.PurchaseForm
	var err error

	result := r.db.Order("id desc").Preload(clause.Associations).Find(&purchaseForm, "car_id = ?", id)
	if result.RowsAffected == 0 {
		err = errors.New("data tidak ditemukan")
	}

	return purchaseForm, err
}
