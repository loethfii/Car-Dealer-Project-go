package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 200
func StatusOk(c *gin.Context, data any, message string) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": message,
	})
}

// // 201
func StatusCreated(c *gin.Context, data any, message string) {
	c.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"message": message,
	})
}

// // 400
func StatusBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}

// // 401
// func StatusForbidden()
//
// // 404
func StatusNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": message,
	})
}

// //serval
func StatusServalInternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": message,
	})

}
