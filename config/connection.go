package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"tugas_akhir_course_net/models"
)

func InitDB() (*gorm.DB, error) {
	config := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/dealer_mobil_gorm?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := config
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Error : ", err)
		panic("Failed to connect DB")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Car{})
	DB.AutoMigrate(&models.SalesPeople{})
	DB.AutoMigrate(&models.PurchaseForm{})

	return DB, err
}
