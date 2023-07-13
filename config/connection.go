package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"tugas_akhir_course_net/models"
)

func InitDB() (*gorm.DB, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Gagal membaca file konfigurasi: %v", err)
	}
	
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbName := viper.GetString("DB_NAME")
	
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
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
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.UploadImage{})
	
	return DB, err
}
