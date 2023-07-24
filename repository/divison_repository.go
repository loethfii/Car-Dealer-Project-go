package repository

import (
	"gorm.io/gorm"
	"tugas_akhir_course_net/models"
)

type DivisionRepository interface {
	FindAll() ([]models.Division, error)
	FindById(id int) (models.Division, error)
	Create(car models.Division) (models.Division, error)
	//Update(id int, updateCar models.CarRequest) (models.CarRequest, error)
	//Delete(id int) (models.Department, error)
}

type divisonRepository struct {
	db *gorm.DB
}

func NewDivisionRepository(db *gorm.DB) *divisonRepository {
	return &divisonRepository{db}
}

func (r *divisonRepository) FindAll() ([]models.Division, error) {
	var divisions []models.Division

	err := r.db.Order("id desc").Find(&divisions).Error

	return divisions, err
}

func (r *divisonRepository) FindById(id int) (models.Division, error) {
	var division models.Division

	err := r.db.Order("id desc").First(&division, "id = ?", id).Error

	return division, err
}

func (r *divisonRepository) Create(division models.Division) (models.Division, error) {
	err := r.db.Create(&division).Error

	return division, err
}
