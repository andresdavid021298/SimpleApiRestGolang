package repository

import (
	"ApiRest/model"
	"ApiRest/util"
)

type MotorcycleRepository struct {
	data []model.Motorcycle
}

func NewMotorcycleRepository() *MotorcycleRepository {
	motorcycles := util.MotorcycleCacheInit()
	return &MotorcycleRepository{data: motorcycles}
}

func (repository *MotorcycleRepository) GetAll() []model.Motorcycle {
	return repository.data
}

func (repository *MotorcycleRepository) GetByBrand(brand string) []model.Motorcycle {
	var motorcycles []model.Motorcycle
	for _, motorcycle := range repository.data {
		if motorcycle.Vehicle.Brand == brand {
			motorcycles = append(motorcycles, motorcycle)
		}
	}
	return motorcycles
}

func (repository *MotorcycleRepository) GetById(id string) *model.Motorcycle {
	for _, motorcycle := range repository.data {
		uuidAsString := motorcycle.Vehicle.Id.String()
		if uuidAsString == id {
			return &motorcycle
		}
	}
	return nil
}

func (repository *MotorcycleRepository) FilterByCC(minCc int, maxCc int) []model.Motorcycle {
	var motorcycles []model.Motorcycle
	for _, motorcycle := range repository.data {
		cc := motorcycle.Cc
		if cc >= minCc && cc <= maxCc {
			motorcycles = append(motorcycles, motorcycle)
		}
	}
	return motorcycles
}

func (repository *MotorcycleRepository) Update(id string, modelUpdated model.Motorcycle) {
	for i, motorcycle := range repository.data {
		uuidAsString := motorcycle.Vehicle.Id.String()
		if uuidAsString == id {
			repository.data[i] = modelUpdated
			break
		}
	}
}

func (repository *MotorcycleRepository) DeleteById(index int) {
	initNewData := repository.data[:index]
	endNewData := repository.data[index+1:]
	repository.data = append(initNewData, endNewData...)
}
