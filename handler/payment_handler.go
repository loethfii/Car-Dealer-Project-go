package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/service"
)

type paymentHandler struct {
	paymentService      service.PaymentService
	purchaseFormService service.PurchaseFormService
}

func NewPaymentHandler(paymentService service.PaymentService, purchaseFormService service.PurchaseFormService) *paymentHandler {
	return &paymentHandler{paymentService, purchaseFormService}
}

func (h paymentHandler) GetPayments(c *gin.Context) {
	payments, _, err := h.paymentService.FindAll()
	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return
	}

	newPayments := []models.PaymentInnerJoinPurchaseForm{}

	for _, v := range payments {
		purchaseForms, _, _, _ := h.purchaseFormService.FindById(v.PurchaseFormId)
		data := helper.ConvToPaymentResponseInnerJoin(v, purchaseForms)

		newPayments = append(newPayments, data)
	}

	length := len(payments)

	message := fmt.Sprintf("%d data ditemukan", length)

	helper.StatusOk(c, newPayments, message)
}

func (h paymentHandler) GetPaymentsId(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	payment, err := h.paymentService.FindById(id)
	if err != nil {
		helper.StatusNotFound(c, "Data tidak ditemukan")
		return
	}

	purchaseForms, _, _, _ := h.purchaseFormService.FindById(payment.PurchaseFormId)

	helper.StatusOk(c, helper.ConvToPaymentResponseInnerJoin(payment, purchaseForms), "Data ditemukan")

}

func (h paymentHandler) PostPayment(c *gin.Context) {

	file, err := c.FormFile("bukti_transfer")
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}
	// Menyimpan file yang diunggah ke direktori uploads
	filename := helper.GenerateFilename(file.Filename)

	err = c.SaveUploadedFile(file, "assets/img/"+filename)

	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return
	}

	stringPurchaseForm := c.PostForm("purchase_form_id")
	purchaseFormId, _ := strconv.Atoi(stringPurchaseForm)

	var newPayment = models.PaymentRequest{
		BuktiTransfer:  filename,
		PurchaseFormId: purchaseFormId,
	}

	payment, err := h.paymentService.Create(newPayment)
	if err != nil {
		if err.Error() == "data purchase form tidak ditemukan" {
			helper.StatusNotFound(c, err.Error())
			return
		} else if err.Error() == "data purchase form id duplikat" {
			helper.StatusServalInternalError(c, err.Error())
			return
		} else {
			helper.StatusServalInternalError(c, err.Error())
		}
	}

	helper.StatusCreated(c, helper.ConvToPaymentResponse(payment), "Data berhasil ditambah")
}

func (h paymentHandler) PutPayment(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	var payment models.PaymentRequest
	err := c.ShouldBind(&payment)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	payment, err = h.paymentService.Update(id, payment)
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
	}

	helper.StatusOk(c, payment, "Data berhasil diupdate")
}

func (h paymentHandler) DeletePayment(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	payment, err := h.paymentService.Delete(id)
	if err != nil {
		helper.StatusNotFound(c, "Data tidak ditemukan")
		return
	}

	helper.StatusOk(c, helper.ConvToPaymentResponse(payment), "Data berhasil terhapus")
}

func (h paymentHandler) ConfirmPayment(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	payment, err := h.paymentService.ConfirmPayment(id)
	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return
	}

	dataID, _ := h.paymentService.FindById(id)
	helper.StatusOk(c, helper.DataConfirmPayment(payment, dataID), "Pembayaran berhasil dikonfirmasi")

}
