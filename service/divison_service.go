package service

import (
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/repository"
)

type DivisionService interface {
	FindAll() ([]models.Division, error)
	FindById(id int) (models.Division, error)
	Create(divisionRequest models.DivisionRequest) (models.Division, error)
	//Update(id int, updateCar models.CarRequest) (models.CarRequest, error)
	//Delete(id int) (models.Car, error)
}

type divisionService struct {
	divisionRepository repository.DivisionRepository
}

func NewDivisonService(divisionRepository repository.DivisionRepository) *divisionService {
	return &divisionService{divisionRepository}
}

func (s *divisionService) FindAll() ([]models.Division, error) {
	divisions, err := s.divisionRepository.FindAll()

	return divisions, err
}

func (s *divisionService) FindById(id int) (models.Division, error) {
	division, err := s.divisionRepository.FindById(id)

	return division, err
}

func (s *divisionService) Create(divisionRequest models.DivisionRequest) (models.Division, error) {
	var newDivision = models.Division{
		NameDivision: divisionRequest.NameDivision,
		Code:         divisionRequest.Code,
		DepartmentID: divisionRequest.DepartmentID,
	}
	division, err := s.divisionRepository.Create(newDivision)

	return division, err
}
