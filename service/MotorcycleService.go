package service

import (
	"ApiRest/model"
	"ApiRest/repository"
	"log"
)

type MotorcycleService struct {
	repository repository.MotorcycleRepository
}

func NewMotorcycleService() *MotorcycleService {
	return &MotorcycleService{*repository.NewMotorcycleRepository()}
}

func (motorcycleService *MotorcycleService) GetAll() []model.Motorcycle {
	log.Println("SEARCHING ALL MOTORCYCLES ")
	return motorcycleService.repository.GetAll()
}

func (motorcycleService *MotorcycleService) GetByBrand(brand string) []model.Motorcycle {
	log.Println("SEARCHING MOTORCYCLES BY BRAND", brand)
	return motorcycleService.repository.GetByBrand(brand)
}

func (motorcycleService *MotorcycleService) GetById(id string) *model.Motorcycle {
	log.Println("SEARCHING MOTORCYCLE BY ID", id)
	return motorcycleService.repository.GetById(id)
}

func (motorcycleService *MotorcycleService) FilterByCC(minCC int, maxCC int) []model.Motorcycle {
	log.Println("SEARCHING MOTORCYCLES BY CC")
	mototcyclesFounded := motorcycleService.repository.FilterByCC(minCC, maxCC)
	log.Println("MOTORCYCLE FOUNDED: ", len(mototcyclesFounded))
	return mototcyclesFounded
}

func (motorcycleService *MotorcycleService) ChangeAmount(id string, newAmount string) *model.Motorcycle {
	log.Println("CHANGING AMOUNT MOTORCYCLES BY ID", id)
	motorcycleForUpdate := motorcycleService.GetById(id)
	if motorcycleForUpdate != nil {
		motorcycleForUpdate.Vehicle.Amount = newAmount
		motorcycleService.repository.Update(id, *motorcycleForUpdate)
	}
	return motorcycleForUpdate
}

func (motorcycleService *MotorcycleService) DeleteMotorcycle(id string) string {
	log.Println("DELETING MOTORCYCLE WITH ID", id)
	index := -1
	var message string
	if len(motorcycleService.GetAll()) == 0 {
		return "Not Motorcycles"
	}
	for i, motorcycle := range motorcycleService.repository.GetAll() {
		uuidAsString := motorcycle.Vehicle.Id.String()
		if uuidAsString == id {
			index = i
			break
		}
	}
	if index != -1 {
		motorcycleService.repository.DeleteById(index)
		message = "Motorcycle deleted successful"
	} else {
		message = "Motorcycle not found"
	}
	return message
}
