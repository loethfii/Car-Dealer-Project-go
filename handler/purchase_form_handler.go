package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/service"
)

type purchaseFormHandler struct {
	purchaseFormService service.PurchaseFormService
	carService          service.CarService
	salesPeopleService  service.SalesPeopleService
}

func NewPurchaseFormHandler(purchaseFormService service.PurchaseFormService, carService service.CarService, salesPeopleService service.SalesPeopleService) *purchaseFormHandler {
	return &purchaseFormHandler{purchaseFormService, carService, salesPeopleService}
}

func (h *purchaseFormHandler) GetPurchaseForm(c *gin.Context) {
	purchaseForms, carId, err := h.purchaseFormService.FindAll()
	if err != nil {
		if err != nil {
			helper.StatusServalInternalError(c, "Terjadi kesalahan internal server.")
			return
		}
	}

	car, err := h.carService.FindById(carId)
	if err != nil {
		if err != nil {
			helper.StatusServalInternalError(c, "Terjadi kesalahan internal server.")
			return
		}
	}

	salesPeople, err := h.salesPeopleService.FindById(3)
	if err != nil {
		if err != nil {
			helper.StatusServalInternalError(c, "Terjadi kesalahan internal server.")
			return
		}
	}

	newPurchaseFormRes := []models.PurchaseFormInnerJoinResponse{}
	for _, val := range purchaseForms {
		data := helper.ConvertToResponseAndInnerJoin(val, car, salesPeople)

		newPurchaseFormRes = append(newPurchaseFormRes, data)
	}

	length := len(newPurchaseFormRes)

	message := fmt.Sprintf("%d data ditemukan", length)

	helper.StatusOk(c, newPurchaseFormRes, message)
}

func (h *purchaseFormHandler) GetPurchaseFormById(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	purchaseForm, err := h.purchaseFormService.FindById(id)
	if err != nil {
		helper.StatusNotFound(c, "Data tidak ditemukan")
		return
	}

	message := fmt.Sprintf("Data ditemukan")

	helper.StatusOk(c, helper.ConvertToReponsePurchaseForm(purchaseForm), message)
}

func (h *purchaseFormHandler) PostPurchaseForm(c *gin.Context) {
	var newPuchaseForm models.PurchaseFormRequest

	if err := c.ShouldBindJSON(&newPuchaseForm); err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	purchaseForm, err := h.purchaseFormService.Create(newPuchaseForm)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	helper.StatusCreated(c, helper.ConvertToReponsePurchaseForm(purchaseForm), "Data berhasil ditambah")

}

func (h *purchaseFormHandler) PutPurchaseForm(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	var newPurchaseForm = models.PurchaseFormRequest{}

	err := c.ShouldBindJSON(&newPurchaseForm)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	purchaseForm, err := h.purchaseFormService.Update(id, newPurchaseForm)
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
	}

	helper.StatusOk(c, purchaseForm, "Berhasil update")
}

func (h *purchaseFormHandler) DeletePurchaseForm(c *gin.Context) {

	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	purchaseForm, err := h.purchaseFormService.Delete(id)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	helper.StatusOk(c, helper.ConvertToReponsePurchaseForm(purchaseForm), "Data Berhasil dihapus.")
}
