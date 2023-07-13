package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tugas_akhir_course_net/config"
	"tugas_akhir_course_net/handler"
	"tugas_akhir_course_net/helper"
	"tugas_akhir_course_net/middlewares"
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
	paymentHandler := handler.NewPaymentHandler(paymentService, purchaseFormService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	fmt.Println(userService)

	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		var car = v1.Group("/car")
		{
			car.GET("/", middlewares.Auth(), carHandler.GetCars)
			car.GET("/:id", middlewares.Auth(), carHandler.GetCarsById)
			car.POST("/post", middlewares.ManagerAndStaffSales(), carHandler.PostCars)
			car.PUT("/put/:id", middlewares.ManagerAndStaffSales(), carHandler.PutCars)
			car.DELETE("/delete/:id", middlewares.ManagerSales(), carHandler.DeleteCars)
		}

		var sales = v1.Group("/sales")
		{
			sales.GET("/", middlewares.ManagerAndStaffSales(), salesPeopleHandler.GetSalesPeople)
			sales.GET("/:id", middlewares.ManagerAndStaffSales(), salesPeopleHandler.GetSalesPeopleById)
			sales.POST("/post", middlewares.ManagerSales(), salesPeopleHandler.PostSalesPeople)
			sales.PUT("/put/:id", middlewares.ManagerAndStaffSales(), salesPeopleHandler.PutSalesPeople)
			sales.DELETE("/delete/:id", middlewares.ManagerSales(), salesPeopleHandler.DeleteSalesPeople)
		}

		var purchaseForm = v1.Group("/purchase-form")
		{
			purchaseForm.GET("/", middlewares.ManagerAndStaffSales(), purchaseFormHandler.GetPurchaseForm)
			purchaseForm.GET("/sales-people/:id", middlewares.ManagerAndStaffSales(), purchaseFormHandler.PurchaseFormFindSalesPepleID)
			purchaseForm.GET("/car/:id", middlewares.ManagerAndStaffSales(), purchaseFormHandler.PurchaseFormFindCarID)
			purchaseForm.GET("/:id", middlewares.ManagerAndStaffSales(), purchaseFormHandler.GetPurchaseFormById)
			purchaseForm.POST("/post", middlewares.Auth(), purchaseFormHandler.PostPurchaseForm)
			purchaseForm.PUT("/put/:id", middlewares.ManagerAndStaffSales(), purchaseFormHandler.PutPurchaseForm)
			purchaseForm.DELETE("/delete/:id", middlewares.ManagerSales(), purchaseFormHandler.DeletePurchaseForm)
		}

		var payment = v1.Group("/payment")
		{
			payment.GET("/", middlewares.ManagerAndStaffSales(), paymentHandler.GetPayments)
			payment.GET("/:id", middlewares.ManagerAndStaffSales(), paymentHandler.GetPaymentsId)
			payment.POST("/post", middlewares.User(), paymentHandler.PostPayment)
			payment.PUT("/put/:id", middlewares.Auth(), paymentHandler.PutPayment)
			payment.DELETE("/delete/:id", middlewares.ManagerSales(), paymentHandler.DeletePayment)
			payment.PATCH("/confirm/:id", middlewares.ManagerSales(), paymentHandler.ConfirmPayment)
		}

		var user = v1.Group("/user")
		{
			user.POST("/register", userHandler.RegisterUser)
			user.POST("/login", userHandler.GenerateToken)
		}
	}

	r.Run()
}
