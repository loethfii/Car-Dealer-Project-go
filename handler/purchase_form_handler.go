package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/service"
)

type purchaseFormHandler struct {
	purchaseFormService service.PurchaseFormService
}

func NewPurchaseFormHandler(purchaseFormService service.PurchaseFormService) *purchaseFormHandler {
	return &purchaseFormHandler{purchaseFormService}
}

func (h *purchaseFormHandler) GetPurchaseForm(c *gin.Context) {
	purchaseForms, err := h.purchaseFormService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newPurchaseFormRes := []models.PurchaseFormResponse{}
	for _, val := range purchaseForms {
		data := helper.ConvertToReponsePurchaseForm(val)

		newPurchaseFormRes = append(newPurchaseFormRes, data)
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusOK,
		"data":   newPurchaseFormRes,
	})

}

func (h *purchaseFormHandler) GetPurchaseFormById(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	purchaseForm, err := h.purchaseFormService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusOK,
		"data":   helper.ConvertToReponsePurchaseForm(purchaseForm),
	})

}

func (h *purchaseFormHandler) PostPurchaseForm(c *gin.Context) {
	var newPuchaseForm models.PurchaseFormRequest

	if err := c.ShouldBindJSON(&newPuchaseForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	purchaseForm, err := h.purchaseFormService.Create(newPuchaseForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusCreated,
		"data":   helper.ConvertToReponsePurchaseForm(purchaseForm),
	})

}

func (h *purchaseFormHandler) PutPurchaseForm(c *gin.Context) {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	var newPurchaseForm = models.PurchaseFormRequest{}

	err := c.ShouldBindJSON(&newPurchaseForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchaseForm, err := h.purchaseFormService.Update(id, newPurchaseForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   purchaseForm,
	})
}

func (h *purchaseFormHandler) DeletePurchaseForm(c *gin.Context) {

	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)

	purchaseForm, err := h.purchaseFormService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   helper.ConvertToReponsePurchaseForm(purchaseForm),
	})
}
