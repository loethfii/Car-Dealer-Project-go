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
	paymentService service.PaymentService
}

func NewPaymentHandler(paymentService service.PaymentService) *paymentHandler {
	return &paymentHandler{paymentService}
}

func (h paymentHandler) GetPayments(c *gin.Context) {
	payments, err := h.paymentService.FindAll()

	var newPayments []models.PaymentResponse

	for _, v := range payments {
		data := models.Payment{
			Id:             v.Id,
			BuktiTransfer:  v.BuktiTransfer,
			IsConfirm:      v.IsConfirm,
			PurchaseFormId: v.PurchaseFormId,
		}

		newPayments = append(newPayments, helper.ConvToPaymentResponse(data))
	}

	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return
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

	helper.StatusOk(c, helper.ConvToPaymentResponse(payment), "Data ditemukan")

}

func (h paymentHandler) PostPayment(c *gin.Context) {
	var newPayment models.PaymentRequest
	err := c.ShouldBind(&newPayment)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}
	payment, err := h.paymentService.Create(newPayment)
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
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
		helper.StatusNotFound(c, "Data tidak ditemukan")
		return
	}

	dataID, _ := h.paymentService.FindById(id)

	helper.StatusOk(c, helper.DataConfirmPayment(payment, dataID), "Data berhasil terhapus")
}
