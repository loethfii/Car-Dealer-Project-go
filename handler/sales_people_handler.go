package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/service"
)

type salesPeopleHandler struct {
	salesPeopleService service.SalesPeopleService
}

func SalesPeopleHandler(salesPeopleService service.SalesPeopleService) *salesPeopleHandler {
	return &salesPeopleHandler{salesPeopleService}
}

func (h *salesPeopleHandler) GetSalesPeople(c *gin.Context) {
	salesPeoples, err := h.salesPeopleService.FindAll()
	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return
	}
	newSalesPeople := []models.SalesPeopleResponse{}
	for _, val := range salesPeoples {
		data := helper.ConvertToResponseSalesPeople(val)

		newSalesPeople = append(newSalesPeople, data)
	}

	length := len(newSalesPeople)

	message := fmt.Sprintf("%d data ditemukan", length)

	helper.StatusOk(c, newSalesPeople, message)
}

func (h *salesPeopleHandler) GetSalesPeopleById(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	salesPeople, err := h.salesPeopleService.FindById(id)
	if err != nil {
		helper.StatusNotFound(c, "Data tidak ditemukan")
		return
	}

	message := fmt.Sprintf("Data ditemukan")

	helper.StatusOk(c, helper.ConvertToResponseSalesPeople(salesPeople), message)
}

func (h *salesPeopleHandler) PostSalesPeople(c *gin.Context) {

	var newSalesPeople = models.SalesPeopleRequest{}

	err := c.ShouldBindJSON(&newSalesPeople)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	salesPeople, err := h.salesPeopleService.Create(newSalesPeople)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	helper.StatusCreated(c, helper.ConvertToResponseSalesPeople(salesPeople), "Data berhasil ditambahkan")
}

func (h *salesPeopleHandler) PutSalesPeople(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	var newSalesPeople = models.SalesPeopleRequest{}

	err := c.ShouldBindJSON(&newSalesPeople)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	salesPople, err := h.salesPeopleService.Update(id, newSalesPeople)
	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return
	}

	helper.StatusOk(c, salesPople, "Berhasil update")
}

func (h *salesPeopleHandler) DeleteSalesPeople(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	salesPeople, err := h.salesPeopleService.Delete(id)
	if err != nil {
		helper.StatusNotFound(c, "Tidak dapat menghapus, data tidak ditemukan!")
		return
	}

	helper.StatusOk(c, helper.ConvertToResponseSalesPeople(salesPeople), "Data Berhasil dihapus.")
}
