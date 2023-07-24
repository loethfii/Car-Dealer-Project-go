package service

import (
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/repository"
)

type DepartmentService interface {
	FindAll() ([]models.DepartementResponse, error)
	FindById(id int) (models.DepartementResponse, error)
	Create(departmentRequest models.DepartmentRequest) (models.Department, error)
	//Update(id int, updateCar models.CarRequest) (models.CarRequest, error)
	//Delete(id int) (models.Car, error)
}

type departmentService struct {
	departmentRepository repository.DepartmentRepository
}

func NewDepartmentService(departmentRepository repository.DepartmentRepository) *departmentService {
	return &departmentService{departmentRepository}
}

func (s *departmentService) FindAll() ([]models.DepartementResponse, error) {
	departments, err := s.departmentRepository.FindAll()

	var departmentresponse []models.DepartementResponse

	for _, each := range departments {
		var divisionsResponse []models.DivisionResponse

		for _, each2 := range each.Divisions {
			dataDivision := models.DivisionResponse{
				ID:           each2.ID,
				NameDivision: each2.NameDivision,
				Code:         each2.Code,
				DepartmentID: each2.DepartmentID,
			}
			divisionsResponse = append(divisionsResponse, dataDivision)
		}

		data := models.DepartementResponse{
			ID:             each.ID,
			NameDepartment: each.NameDepartment,
			Code:           each.Code,
			Divisions:      divisionsResponse,
		}

		departmentresponse = append(departmentresponse, data)
	}

	return departmentresponse, err
}

func (s *departmentService) FindById(id int) (models.DepartementResponse, error) {
	department, err := s.departmentRepository.FindById(id)

	var divisionResponse []models.DivisionResponse

	for _, each := range department.Divisions {
		divisionRes := models.DivisionResponse{
			ID:           each.ID,
			NameDivision: each.NameDivision,
			Code:         each.Code,
			DepartmentID: each.DepartmentID,
		}
		divisionResponse = append(divisionResponse, divisionRes)
	}

	var departmentsResponse = models.DepartementResponse{
		ID:             department.ID,
		NameDepartment: department.NameDepartment,
		Code:           department.Code,
		Divisions:      divisionResponse,
	}

	return departmentsResponse, err
}

func (s *departmentService) Create(departmentRequest models.DepartmentRequest) (models.Department, error) {

	var newDepartment = models.Department{
		NameDepartment: departmentRequest.NameDepartment,
		Code:           departmentRequest.Code,
	}
	department, err := s.departmentRepository.Create(newDepartment)

	return department, err
}
