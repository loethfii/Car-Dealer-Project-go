package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var reqUser models.UserRequest

	err := c.ShouldBindJSON(&reqUser)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	user, err := h.userService.RegisterUser(reqUser)
	if err != nil {
		helper.StatusBadRequest(c, err.Error())
		return
	}

	helper.StatusCreated(c, user, "Register Berhasil")
}

func (h *userHandler) GenerateToken(c *gin.Context) {

	var reqToken models.TokenRequest

	err := c.ShouldBindJSON(&reqToken)
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
	}

	token, err := h.userService.LoginUser(reqToken)
	if err != nil {
		helper.StatusBadRequest(c, "UnAuthrizied")
		return
	}

	err = h.userService.CredentialEmailPassword(reqToken)
	if err != nil {
		helper.StatusServalInternalError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
