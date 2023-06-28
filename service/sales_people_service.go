package service

import (
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/repository"
)

type SalesPeopleService interface {
	FindAll() ([]models.SalesPeople, error)
	FindById(id int) (models.SalesPeople, error)
	Create(salesPeople models.SalesPeopleRequest) (models.SalesPeople, error)
	Update(id int, updateSalesPeople models.SalesPeopleRequest) (models.SalesPeopleRequest, error)
	Delete(id int) (models.SalesPeople, error)
}

type salesPeopleService struct {
	salesPeopleRepository repository.SalesPeopleRepository
}

func NewSalesPeopleService(salesPeopleRepository repository.SalesPeopleRepository) *salesPeopleService {
	return &salesPeopleService{salesPeopleRepository}
}

func (s *salesPeopleService) FindAll() ([]models.SalesPeople, error) {
	salesPeoples, err := s.salesPeopleRepository.FindAll()
	return salesPeoples, err
}

func (s *salesPeopleService) FindById(id int) (models.SalesPeople, error) {
	salesPeople, err := s.salesPeopleRepository.FindById(id)
	return salesPeople, err
}

func (s *salesPeopleService) Create(salesPeople models.SalesPeopleRequest) (models.SalesPeople, error) {
	var newSalesPeople = models.SalesPeople{
		NamaSales:   salesPeople.NamaSales,
		Nip:         salesPeople.Nip,
		NomerTelpon: salesPeople.NomerTelpon,
		Bagian:      salesPeople.Bagian,
	}

	newSalesPeople, err := s.salesPeopleRepository.Create(newSalesPeople)
	return newSalesPeople, err
}

func (s *salesPeopleService) Update(id int, updateSalesPeople models.SalesPeopleRequest) (models.SalesPeopleRequest, error) {
	updateSalesPeople, err := s.salesPeopleRepository.Update(id, updateSalesPeople)
	return updateSalesPeople, err
}

func (s *salesPeopleService) Delete(id int) (models.SalesPeople, error) {
	salesPeople, err := s.salesPeopleRepository.Delete(id)

	return salesPeople, err
}
