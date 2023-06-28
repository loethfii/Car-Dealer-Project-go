package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	helper.Error(err)
	newSalesPeople := []models.SalesPeopleResponse{}
	for _, val := range salesPeoples {
		data := helper.ConvertToResponseSalesPeople(val)

		newSalesPeople = append(newSalesPeople, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   newSalesPeople,
	})
}

func (h *salesPeopleHandler) GetSalesPeopleById(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	salesPeople, err := h.salesPeopleService.FindById(id)
	helper.Error(err)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   helper.ConvertToResponseSalesPeople(salesPeople),
	})
}

func (h *salesPeopleHandler) PostSalesPeople(c *gin.Context) {

	var newSalesPeople = models.SalesPeopleRequest{}

	err := c.ShouldBindJSON(&newSalesPeople)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesPeople, err := h.salesPeopleService.Create(newSalesPeople)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   helper.ConvertToResponseSalesPeople(salesPeople),
	})
}

func (h *salesPeopleHandler) PutSalesPeople(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	var newSalesPeople = models.SalesPeopleRequest{}

	err := c.ShouldBindJSON(&newSalesPeople)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesPople, err := h.salesPeopleService.Update(id, newSalesPeople)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   salesPople,
	})
}

func (h *salesPeopleHandler) DeleteSalesPeople(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	salesPeople, err := h.salesPeopleService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   helper.ConvertToResponseSalesPeople(salesPeople),
	})
}

func (h *carHandler) D(c *gin.Context) {

	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	car, err := h.carService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   helper.ConvertToResponseCar(car),
	})
}
