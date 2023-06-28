package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/service"
)

type carHandler struct {
	carService service.CarService
}

func NewCarHandler(carService service.CarService) *carHandler {
	return &carHandler{carService}
}

func (h *carHandler) GetCars(c *gin.Context) {
	cars, err := h.carService.FindAll()
	helper.Error(err)
	newCars := []models.CarResponse{}
	for _, val := range cars {
		data := helper.ConvertToResponseCar(val)

		newCars = append(newCars, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   newCars,
	})
}

func (h *carHandler) GetCarsById(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	car, err := h.carService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak ditemukan.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   helper.ConvertToResponseCar(car),
	})
}

func (h *carHandler) PostCars(c *gin.Context) {

	var newCar = models.CarRequest{}

	err := c.ShouldBindJSON(&newCar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car, err := h.carService.Create(newCar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   helper.ConvertToResponseCar(car),
	})
}

func (h *carHandler) PutCars(c *gin.Context) {

	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	var newCar = models.CarRequest{}

	err := c.ShouldBindJSON(&newCar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car, err := h.carService.Update(id, newCar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   car,
	})
}

func (h *carHandler) DeleteCars(c *gin.Context) {

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
