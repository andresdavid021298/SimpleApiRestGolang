package repository

import (
	"ApiRest/model"
	"ApiRest/util"
)

type CarRepository struct {
	data []model.Car
}

func NewCarRepository() *CarRepository {
	cars := util.CarCacheInit()
	return &CarRepository{data: cars}
}

func (repository *CarRepository) GetAll() []model.Car {
	return repository.data
}

func (repository *CarRepository) GetByBrand(brand string) []model.Car {
	var cars []model.Car
	for _, car := range repository.data {
		if car.Vehicle.Brand == brand {
			cars = append(cars, car)
		}
	}
	return cars
}

func (repository *CarRepository) GetUsed() []model.Car {
	var cars []model.Car
	for _, car := range repository.data {
		if car.IsUsed {
			cars = append(cars, car)
		}
	}
	return cars
}

func (repository *CarRepository) Insert(newCar model.Car) {
	repository.data = append(repository.data, newCar)
}

func (repository *CarRepository) DeleteById(index int) {
	initNewData := repository.data[:index]
	endNewData := repository.data[index+1:]
	repository.data = append(initNewData, endNewData...)
}
