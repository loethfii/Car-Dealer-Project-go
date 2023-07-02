package service

import (
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/repository"
)

type PaymentService interface {
	FindAll() ([]models.Payment, error)
	FindById(id int) (models.Payment, error)
	Create(payment models.PaymentRequest) (models.Payment, error)
	Update(id int, updatePayment models.PaymentRequest) (models.PaymentRequest, error)
	Delete(id int) (models.Payment, error)
	ConfirmPayment(id int) (models.PaymentRequest, error)
}

type paymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) *paymentService {
	return &paymentService{paymentRepo}
}

func (s *paymentService) FindAll() ([]models.Payment, error) {
	payments, err := s.paymentRepository.FindAll()
	return payments, err
}

func (s *paymentService) FindById(id int) (models.Payment, error) {
	payment, err := s.paymentRepository.FindById(id)

	return payment, err
}

func (s *paymentService) Create(payment models.PaymentRequest) (models.Payment, error) {

	var newPayment = models.Payment{
		BuktiTransfer:  payment.BuktiTransfer,
		IsConfirm:      payment.IsConfirm,
		PurchaseFormId: payment.PurchaseFormId,
	}

	newPayment, err := s.paymentRepository.Create(newPayment)
	return newPayment, err
}

func (s *paymentService) Update(id int, updatePayment models.PaymentRequest) (models.PaymentRequest, error) {
	updatePayment, err := s.paymentRepository.Update(id, updatePayment)
	return updatePayment, err
}

func (s *paymentService) Delete(id int) (models.Payment, error) {
	payment, err := s.paymentRepository.Delete(id)

	return payment, err
}

func (s *paymentService) ConfirmPayment(id int) (models.PaymentRequest, error) {

	var confirmStatus models.PaymentRequest

	confirmStatus.IsConfirm = true

	payment, err := s.paymentRepository.ConfirmPayment(id, confirmStatus)

	return payment, err
}
