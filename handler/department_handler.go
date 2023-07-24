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

type departmentHandler struct {
	departmentService service.DepartmentService
}

func NewDepartmentHandler(departmentService service.DepartmentService) *departmentHandler {
	return &departmentHandler{departmentService}
}

func (h *departmentHandler) GetDepartments(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"Message": "Not Allowed",
		})
		return
	} else {
		//berhasil

		departments, err := h.departmentService.FindAll()
		if err != nil {
			helper.StatusServalInternalError(c, err.Error())
			return
		}

		message := fmt.Sprintf("%d data ditemukan", len(departments))

		helper.StatusOk(c, departments, message)
	}
}

func (h *departmentHandler) PostDepartment(c *gin.Context) {
	var departmentRequest models.DepartmentRequest

	err := c.ShouldBindJSON(&departmentRequest)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	department, err := h.departmentService.Create(departmentRequest)
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
	}

	helper.StatusCreated(c, department, "Data berhasil ditambah")
}

func (h *departmentHandler) GetDepartmentByID(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	department, err := h.departmentService.FindById(id)
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
	}

	message := fmt.Sprintf("data ditemukan")

	helper.StatusOk(c, department, message)
}
