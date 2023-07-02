package repository

import (
	"gorm.io/gorm"
	"tugas_akhir_course_net/models"
)

type PaymentRepository interface {
	FindAll() ([]models.Payment, error)
	FindById(id int) (models.Payment, error)
	Create(payment models.Payment) (models.Payment, error)
	Update(id int, updatePayment models.PaymentRequest) (models.PaymentRequest, error)
	Delete(id int) (models.Payment, error)
	ConfirmPayment(id int, confirm models.PaymentRequest) (models.PaymentRequest, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) FindAll() ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) FindById(id int) (models.Payment, error) {
	var payment models.Payment

	err := r.db.First(&payment, id).Error

	return payment, err
}

func (r *paymentRepository) Create(payment models.Payment) (models.Payment, error) {
	err := r.db.Create(&payment).Error

	return payment, err
}

func (r *paymentRepository) Update(id int, updatePayment models.PaymentRequest) (models.PaymentRequest, error) {
	err := r.db.Model(&models.Payment{}).Where("id = ?", id).Updates(&updatePayment).Error
	r.db.Updates(&updatePayment)

	return updatePayment, err
}

func (r *paymentRepository) Delete(id int) (models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, id).Error
	r.db.Delete(&payment)

	return payment, err
}

func (r *paymentRepository) ConfirmPayment(id int, confirm models.PaymentRequest) (models.PaymentRequest, error) {
	err := r.db.Model(&models.Payment{}).Where("id = ?", id).Updates(&confirm).Error
	r.db.Updates(&confirm)

	return confirm, err
}
