package service

import (
	"ApiRest/model"
	"ApiRest/repository"
	"ApiRest/request"
	"log"
)

type CarService struct {
	repository repository.CarRepository
}

func NewCarService() *CarService {
	return &CarService{*repository.NewCarRepository()}
}

func (carService *CarService) GetAll() []model.Car {
	log.Println("SEARCHING ALL CARS ")
	return carService.repository.GetAll()
}

func (carService *CarService) GetByBrand(brand string) []model.Car {
	log.Println("SEARCHING CARS BY BRAND", brand)
	return carService.repository.GetByBrand(brand)
}

func (carService *CarService) GetUsed() []model.Car {
	log.Println("SEARCHING USED CARS")
	return carService.repository.GetUsed()
}

func (carService *CarService) InsertNewCar(request request.NewCarRequest) string {
	log.Println("INSERTING NEW CAR")
	km := *request.Km
	if isUsed := *request.IsUsed; !isUsed {
		km = 0
	}
	newCar := model.NewCar(request.Brand, request.Model, request.Amount, request.Color, *request.IsUsed, km)
	carService.repository.Insert(newCar)
	return "Ok!"
}

func (carService *CarService) DeleteCar(id string) string {
	log.Println("DELETING CAR WITH ID", id)
	index := -1
	var message string
	if len(carService.GetAll()) == 0 {
		return "Not cars"
	}
	for i, car := range carService.repository.GetAll() {
		uuidAsString := car.Vehicle.Id.String()
		if uuidAsString == id {
			index = i
			break
		}
	}
	if index != -1 {
		carService.repository.DeleteById(index)
		message = "Car deleted successful"
	} else {
		message = "Car not found"
	}
	return message
}
