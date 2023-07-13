package handler

import (
	"fmt"
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
	if err != nil {
		helper.StatusServalInternalError(c, "Terjadi kesalahan internal server.")
		return
	}
	newCars := []models.CarResponse{}
	for _, val := range cars {
		data := helper.ConvertToResponseCar(val)

		newCars = append(newCars, data)
	}

	length := len(newCars)

	message := fmt.Sprintf("%d data ditemukan", length)

	helper.StatusOk(c, newCars, message)
}

func (h *carHandler) GetCarsById(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	car, err := h.carService.FindById(id)
	if err != nil {
		helper.StatusNotFound(c, "Data tidak ditemukan")
		return
	}

	message := fmt.Sprintf("Data ditemukan")

	helper.StatusOk(c, helper.ConvertToResponseCar(car), message)
}

func (h *carHandler) PostCars(c *gin.Context) {

	var newCar = models.CarRequest{}

	err := c.ShouldBindJSON(&newCar)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	car, err := h.carService.Create(newCar)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	helper.StatusCreated(c, helper.ConvertToRequestCar(car), "Data berhasil ditambahkan")
}

func (h *carHandler) PutCars(c *gin.Context) {

	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	var newCar = models.CarRequest{}

	err := c.ShouldBindJSON(&newCar)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	car, err := h.carService.Update(id, newCar)
	if err != nil {
		helper.StatusNotFound(c, err.Error())
		return
	}

	fmt.Sprintln(car)

	c.JSON(http.StatusOK, gin.H{
		"Message": "berhasil update",
	})
}

func (h *carHandler) DeleteCars(c *gin.Context) {

	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	car, err := h.carService.Delete(id)
	if err != nil {
		helper.StatusNotFound(c, "Tidak dapat menghapus, data tidak ditemukan!")
		return
	}

	helper.StatusOk(c, helper.ConvertToResponseCar(car), "Data Berhasil dihapus.")
}
