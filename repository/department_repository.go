package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tugas_akhir_course_net/models"
)

type DepartmentRepository interface {
	FindAll() ([]models.Department, error)
	FindById(id int) (models.Department, error)
	Create(department models.Department) (models.Department, error)
	Update(id int, updateDepartment models.DepartmentRequest) (models.DepartmentRequest, error)
	Delete(id int) (models.Department, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *departmentRepository {
	return &departmentRepository{db}
}

func (r *departmentRepository) FindAll() ([]models.Department, error) {
	var departments []models.Department

	err := r.db.Order("id desc").Preload(clause.Associations).Find(&departments).Error

	return departments, err
}

func (r *departmentRepository) FindById(id int) (models.Department, error) {
	var department models.Department

	err := r.db.Order("id desc").Preload(clause.Associations).First(&department, "id = ?", id).Error

	return department, err
}

func (r *departmentRepository) Create(department models.Department) (models.Department, error) {
	err := r.db.Create(&department).Error

	return department, err
}

func (r *departmentRepository) Update(id int, updateDepartment models.DepartmentRequest) (models.DepartmentRequest, error) {
	var department models.Department
	err := r.db.Model(&department).Where("id ?", id).Updates(&updateDepartment).Error

	return updateDepartment, err
}

func (r *departmentRepository) Delete(id int) (models.Department, error) {
	var department models.Department
	err := r.db.Delete(department, "id ?", id).Error

	return department, err
}
