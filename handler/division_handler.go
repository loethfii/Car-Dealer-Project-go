package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/service"
)

type divisionHandler struct {
	devisionService service.DivisionService
}

func NewDivisionHandler(devisionService service.DivisionService) *divisionHandler {
	return &divisionHandler{devisionService}
}

func (h divisionHandler) GetDevisions(c *gin.Context) {
	divisions, err := h.devisionService.FindAll()
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
	}

	message := fmt.Sprintf("%d Data ditemukan", len(divisions))

	helper.StatusOk(c, divisions, message)
}

func (h divisionHandler) CreateDivision(c *gin.Context) {
	var devisionRequest models.DivisionRequest
	err := c.ShouldBindJSON(&devisionRequest)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	division, err := h.devisionService.Create(devisionRequest)
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
	}

	helper.StatusCreated(c, division, "Data berhasil dibuat")
}
