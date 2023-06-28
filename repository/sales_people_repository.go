package repository

import (
	"gorm.io/gorm"
	"tugas_akhir_course_net/models"
)

type SalesPeopleRepository interface {
	FindAll() ([]models.SalesPeople, error)
	FindById(id int) (models.SalesPeople, error)
	Create(car models.SalesPeople) (models.SalesPeople, error)
	Update(id int, updateCar models.SalesPeopleRequest) (models.SalesPeopleRequest, error)
	Delete(id int) (models.SalesPeople, error)
}

type salesPeopleRepository struct {
	db *gorm.DB
}

func NewSalesPeopleRepository(db *gorm.DB) *salesPeopleRepository {
	return &salesPeopleRepository{db}
}

func (r *salesPeopleRepository) FindAll() ([]models.SalesPeople, error) {
	var salesPeoples []models.SalesPeople

	err := r.db.Order("id desc").Find(&salesPeoples).Error
	r.db.Order("id desc").Preload("PurchaseForms", func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc")
	}).Find(&salesPeoples)

	return salesPeoples, err
}

func (r *salesPeopleRepository) FindById(id int) (models.SalesPeople, error) {
	var salesPeople models.SalesPeople

	err := r.db.First(&salesPeople, id).Error
	r.db.Preload("PurchaseForms", func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc")
	}).Find(&salesPeople)

	return salesPeople, err
}

func (r *salesPeopleRepository) Create(salesPeople models.SalesPeople) (models.SalesPeople, error) {
	err := r.db.Create(&salesPeople).Error

	return salesPeople, err
}

func (r *salesPeopleRepository) Update(id int, updatesalesPeople models.SalesPeopleRequest) (models.SalesPeopleRequest, error) {

	err := r.db.Model(&models.SalesPeople{}).Where("id = ?", id).Updates(&updatesalesPeople).Error

	return updatesalesPeople, err
}

func (r *salesPeopleRepository) Delete(id int) (models.SalesPeople, error) {

	var salesPeople models.SalesPeople

	r.db.First(&salesPeople, id)

	err := r.db.Delete(&salesPeople).Error

	return salesPeople, err
}
