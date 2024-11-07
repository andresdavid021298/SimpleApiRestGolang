package model

import "github.com/google/uuid"

type Car struct {
	Vehicle Vehicle
	IsUsed  bool
	Km      int
}

func NewCar(brand string, model string, amount string, color string, isUsed bool, km int) Car {
	uuidGenerated := uuid.New()
	basedVehicle := Vehicle{Id: uuidGenerated, Amount: amount, Color: color, Brand: brand, Model: model}
	car := Car{IsUsed: isUsed, Km: km, Vehicle: basedVehicle}
	return car
}
