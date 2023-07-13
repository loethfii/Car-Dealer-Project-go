package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tugas_akhir_course_net/models"
)

type PaymentRepository interface {
	FindAll() ([]models.Payment, int, error)
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

func (r *paymentRepository) FindAll() ([]models.Payment, int, error) {
	var payments []models.Payment
	err := r.db.Order("id desc").Preload(clause.Associations).Find(&payments).Error

	var purchaseFormId int

	for _, each := range payments {
		purchaseFormId = each.PurchaseFormId
	}
	return payments, purchaseFormId, err
}

func (r *paymentRepository) FindById(id int) (models.Payment, error) {
	var payment models.Payment

	err := r.db.First(&payment, id).Error

	return payment, err
}

func (r *paymentRepository) Create(payment models.Payment) (models.Payment, error) {
	var err error

	var purchaseForm models.PurchaseForm
	purchaseFormID := payment.PurchaseFormId

	result := r.db.First(&purchaseForm, purchaseFormID)

	if result.RowsAffected == 0 {
		err = errors.New("data purchase form tidak ditemukan")
	} else {
		errDB := r.db.Create(&payment).Error
		if errDB != nil {
			err = errors.New("data purchase form id duplikat")
		}
	}

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

	var payment models.Payment

	result := r.db.First(&payment, "id = ? ", id)
	if result.Error != nil {
		return confirm, errors.New("Data tidak ditemukan")
	}

	purchaseFormID := payment.PurchaseFormId

	var purchaseForm models.PurchaseForm

	r.db.First(&purchaseForm, "id = ? ", purchaseFormID)

	if payment.IsConfirm == false {
		if purchaseForm.HarusInden == true {
			err := r.db.Model(&models.Payment{}).Where("id = ?", id).Updates(&confirm).Error
			return confirm, err
		} else {
			err := r.db.Model(&models.Payment{}).Where("id = ?", id).Updates(&confirm).Error

			carID := purchaseForm.CarId

			r.db.Model(&models.Car{}).Where("id = ?", carID).Update("qty", gorm.Expr("qty - ?", 1))

			return confirm, err
		}
	} else {
		return confirm, errors.New("Pembayaran ini sudah dikonfirmasi tidak dapat dikonfirmasi ulang")
	}

}
