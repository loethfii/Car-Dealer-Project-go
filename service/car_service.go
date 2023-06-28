package service

import (
	"tugas_akhir_course_net/models"
	"tugas_akhir_course_net/repository"
)

type CarService interface {
	FindAll() ([]models.Car, error)
	FindById(id int) (models.Car, error)
	Create(car models.CarRequest) (models.Car, error)
	Update(id int, updateCar models.CarRequest) (models.CarRequest, error)
	Delete(id int) (models.Car, error)
}

type carService struct {
	carRepo repository.CarRepository
}

func NewCarService(carRepo repository.CarRepository) *carService {
	return &carService{carRepo}
}

func (s *carService) FindAll() ([]models.Car, error) {
	cars, err := s.carRepo.FindAll()
	return cars, err
}

func (s *carService) FindById(id int) (models.Car, error) {
	car, err := s.carRepo.FindById(id)
	return car, err
}

func (s *carService) Create(car models.CarRequest) (models.Car, error) {
	var newCar = models.Car{
		NamaMobil:    car.NamaMobil,
		TipeMobil:    car.TipeMobil,
		JenisMobil:   car.JenisMobil,
		BahanBakar:   car.BahanBakar,
		Isi_Silinder: car.Isi_Silinder,
		Warna:        car.Warna,
		Transmisi:    car.Transmisi,
		Harga:        car.Harga,
		Qty:          car.Qty,
	}

	newCar, err := s.carRepo.Create(newCar)
	return newCar, err
}

func (s *carService) Update(id int, updateCar models.CarRequest) (models.CarRequest, error) {
	updateCar, err := s.carRepo.Update(id, updateCar)
	return updateCar, err
}

func (s *carService) Delete(id int) (models.Car, error) {
	car, err := s.carRepo.Delete(id)

	return car, err
}
