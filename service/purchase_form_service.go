package service

import (
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/repository"
)

type PurchaseFormService interface {
	FindAll() ([]models.PurchaseForm, error)
	FindById(id int) (models.PurchaseForm, error)
	Create(car models.PurchaseFormRequest) (models.PurchaseForm, error)
	Update(id int, updateCar models.PurchaseFormRequest) (models.PurchaseFormRequest, error)
	Delete(id int) (models.PurchaseForm, error)
}

type purchaseFormService struct {
	purchaseFormRepository repository.PurchaseFormRepository
}

func NewPurchaseFormService(purchaseFormRepository repository.PurchaseFormRepository) *purchaseFormService {
	return &purchaseFormService{purchaseFormRepository}
}

func (s *purchaseFormService) FindAll() ([]models.PurchaseForm, error) {
	purchaseForms, err := s.purchaseFormRepository.FindAll()
	return purchaseForms, err
}

func (s *purchaseFormService) FindById(id int) (models.PurchaseForm, error) {
	purchaseForm, err := s.purchaseFormRepository.FindById(id)
	return purchaseForm, err
}

func (s *purchaseFormService) Create(puchaseForm models.PurchaseFormRequest) (models.PurchaseForm, error) {
	var newPurchaseForm = models.PurchaseForm{
		NamaLengkapPembeli: puchaseForm.NamaLengkapPembeli,
		NomerKTP:           puchaseForm.NomerKTP,
		AlamatRumah:        puchaseForm.AlamatRumah,
		NomerDebit:         puchaseForm.NomerDebit,
		CarId:              puchaseForm.CarId,
		HarusInden:         puchaseForm.HarusInden,
		LamaInden:          puchaseForm.LamaInden,
		CustomPlat:         puchaseForm.CustomPlat,
		TambahanKit:        puchaseForm.TambahanKit,
		SalesPeopleId:      puchaseForm.SalesPeopleId,
	}

	purchaseForm, err := s.purchaseFormRepository.Create(newPurchaseForm, puchaseForm.CarId, puchaseForm.HarusInden)
	return purchaseForm, err
}

func (s *purchaseFormService) Update(id int, updatePurchaseForm models.PurchaseFormRequest) (models.PurchaseFormRequest, error) {
	updatePurchaseForm, err := s.purchaseFormRepository.Update(id, updatePurchaseForm)
	return updatePurchaseForm, err
}

func (s *purchaseFormService) Delete(id int) (models.PurchaseForm, error) {
	car, err := s.purchaseFormRepository.Delete(id)

	return car, err
}
