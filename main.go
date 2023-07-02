package main

import (
	"github.com/gin-gonic/gin"
	"tugas_akhir_course_net/config"
	"tugas_akhir_course_net/handler"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/repository"
	"tugas_akhir_course_net/service"
)

func main() {
	db, err := config.InitDB()
	helper.Error(err)

	carRepository := repository.NewCarRepository(db)
	carService := service.NewCarService(carRepository)
	carHandler := handler.NewCarHandler(carService)

	salesPeopleRepository := repository.NewSalesPeopleRepository(db)
	salesPeopleService := service.NewSalesPeopleService(salesPeopleRepository)
	salesPeopleHandler := handler.SalesPeopleHandler(salesPeopleService)

	purchaseFormRepository := repository.NewPurchaseFormRepository(db)
	purchaseFormService := service.NewPurchaseFormService(purchaseFormRepository)
	purchaseFormHandler := handler.NewPurchaseFormHandler(purchaseFormService, carService, salesPeopleService)

	paymentRepository := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	r := gin.Default()

	r.GET("/cars", carHandler.GetCars)
	r.GET("/cars/:id", carHandler.GetCarsById)
	r.POST("/cars/post", carHandler.PostCars)
	r.PUT("/cars/put/:id", carHandler.PutCars)
	r.DELETE("/cars/delete/:id", carHandler.DeleteCars)

	r.GET("/sales", salesPeopleHandler.GetSalesPeople)
	r.GET("/sales/:id", salesPeopleHandler.GetSalesPeopleById)
	r.POST("/sales/post", salesPeopleHandler.PostSalesPeople)
	r.PUT("/sales/put/:id", salesPeopleHandler.PutSalesPeople)
	r.DELETE("/sales/delete/:id", salesPeopleHandler.DeleteSalesPeople)

	r.GET("/purchase-forms", purchaseFormHandler.GetPurchaseForm)
	r.GET("/purchase-forms/:id", purchaseFormHandler.GetPurchaseFormById)
	r.POST("/purchase-forms/post", purchaseFormHandler.PostPurchaseForm)
	r.PUT("/purchase-forms/put/:id", purchaseFormHandler.PutPurchaseForm)
	r.DELETE("/purchase-forms/delete/:id", purchaseFormHandler.DeletePurchaseForm)

	r.GET("/payments", paymentHandler.GetPayments)
	r.GET("/payments/:id", paymentHandler.GetPaymentsId)
	r.POST("/payments/post", paymentHandler.PostPayment)
	r.PUT("/payments/put/:id", paymentHandler.PutPayment)
	r.DELETE("/payments/delete/:id", paymentHandler.DeletePayment)
	r.PATCH("/payments/confirm/:id", paymentHandler.ConfirmPayment)

	r.Run()
}
