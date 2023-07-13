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
	purchaseForms, err := h.purchaseFormService.FindAll()
	if err != nil {
		if err != nil {
			helper.StatusServalInternalError(c, "Terjadi kesalahan internal server.")
			return
		}
	}

	newPurchaseFormRes := []models.PurchaseFormInnerJoinResponse{}
	for _, val := range purchaseForms {
		car, _ := h.carService.FindById(val.CarId)
		salesPeople, _ := h.salesPeopleService.FindById(val.SalesPeopleId)

		//data := models.PurchaseFormInnerJoinResponse{}

		data := helper.ConvertFromPurchaseFormToPurchaseFormResponse(val, car, salesPeople)

		newPurchaseFormRes = append(newPurchaseFormRes, data)
	}

	length := len(newPurchaseFormRes)

	message := fmt.Sprintf("%d data ditemukan", length)

	helper.StatusOk(c, newPurchaseFormRes, message)
}

func (h *purchaseFormHandler) GetPurchaseFormById(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	purchaseForm, carId, salesPeopleId, err := h.purchaseFormService.FindById(id)
	if err != nil {
		helper.StatusNotFound(c, "Data tidak ditemukan")
		return
	}

	car, err := h.carService.FindById(carId)
	if err != nil {
		if err != nil {
			helper.StatusServalInternalError(c, "Terjadi kesalahan internal server.")
			return
		}
	}

	salesPeople, err := h.salesPeopleService.FindById(salesPeopleId)
	if err != nil {
		if err != nil {
			helper.StatusServalInternalError(c, "Terjadi kesalahan internal server.")
			return
		}
	}

	message := fmt.Sprintf("Data ditemukan")

	//helper.StatusOk(c, helper.ConvertToReponsePurchaseForm(purchaseForm), message)
	helper.StatusOk(c, helper.ConvertFromPurchaseFormToPurchaseFormResponse(purchaseForm, car, salesPeople), message)

}

func (h *purchaseFormHandler) PostPurchaseForm(c *gin.Context) {
	var newPuchaseForm models.PurchaseFormRequest

	if err := c.ShouldBindJSON(&newPuchaseForm); err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	purchaseForm, err := h.purchaseFormService.Create(newPuchaseForm)
	if err != nil {
		helper.StatusNotFound(c, err.Error())
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
		helper.StatusNotFound(c, err.Error())
		return
	}

	helper.StatusOk(c, purchaseForm, "Berhasil update")
}

func (h *purchaseFormHandler) DeletePurchaseForm(c *gin.Context) {

	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	purchaseForm, err := h.purchaseFormService.Delete(id)
	if err != nil {
		helper.StatusNotFound(c, "Tidak dapat menghapus, data tidak ditemukan")
		return
	}

	helper.StatusOk(c, helper.ConvertToReponsePurchaseForm(purchaseForm), "Data Berhasil dihapus.")
}

func (h *purchaseFormHandler) PurchaseFormFindSalesPepleID(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	purchaseFormsFindSalesPeopleID, err := h.purchaseFormService.FindBySalesPeopleID(id)
	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return

	}

	newPurchaseFormRes := []models.PurchaseFormInnerJoinResponse{}

	for _, val := range purchaseFormsFindSalesPeopleID {
		car, _ := h.carService.FindById(val.CarId)
		salesPeople, _ := h.salesPeopleService.FindById(val.SalesPeopleId)

		data := helper.ConvertFromPurchaseFormToPurchaseFormResponse(val, car, salesPeople)

		newPurchaseFormRes = append(newPurchaseFormRes, data)
	}

	length := len(newPurchaseFormRes)

	message := fmt.Sprintf("%d Data ditemukan", length)

	helper.StatusOk(c, newPurchaseFormRes, message)

}

func (h *purchaseFormHandler) PurchaseFormFindCarID(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	purchaseFormsFindSalesPeopleID, err := h.purchaseFormService.FindByCarID(id)
	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return
	}

	newPurchaseFormRes := []models.PurchaseFormInnerJoinResponse{}

	for _, val := range purchaseFormsFindSalesPeopleID {
		car, _ := h.carService.FindById(val.CarId)
		salesPeople, _ := h.salesPeopleService.FindById(val.SalesPeopleId)

		data := helper.ConvertFromPurchaseFormToPurchaseFormResponse(val, car, salesPeople)

		newPurchaseFormRes = append(newPurchaseFormRes, data)
	}

	length := len(newPurchaseFormRes)

	message := fmt.Sprintf("%d Data ditemukan", length)

	helper.StatusOk(c, newPurchaseFormRes, message)
}
