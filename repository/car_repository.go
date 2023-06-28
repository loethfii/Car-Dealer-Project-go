package repository

import (
	"gorm.io/gorm"
	"tugas_akhir_course_net/models"
)

type CarRepository interface {
	FindAll() ([]models.Car, error)
	FindById(id int) (models.Car, error)
	Create(car models.Car) (models.Car, error)
	Update(id int, updateCar models.CarRequest) (models.CarRequest, error)
	Delete(id int) (models.Car, error)
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *carRepository {
	return &carRepository{db}
}

func (r *carRepository) FindAll() ([]models.Car, error) {
	var cars []models.Car

	var err error

	err = r.db.Order("id desc").Find(&cars).Error
	err = r.db.Order("id desc").Preload("PurchaseForms", func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc")
	}).Find(&cars).Error

	return cars, err
}

func (r *carRepository) FindById(id int) (models.Car, error) {
	var car models.Car

	err := r.db.First(&car, id).Error
	r.db.Preload("PurchaseForms", func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc")
	}).Find(&car)

	return car, err
}

func (r *carRepository) Create(car models.Car) (models.Car, error) {
	err := r.db.Create(&car).Error

	return car, err
}

func (r *carRepository) Update(id int, updateCar models.CarRequest) (models.CarRequest, error) {

	err := r.db.Model(&models.Car{}).Where("id = ?", id).Updates(&updateCar).Error

	return updateCar, err
}

func (r *carRepository) Delete(id int) (models.Car, error) {

	var car models.Car

	r.db.First(&car, id)

	err := r.db.Delete(&car).Error

	return car, err
}
