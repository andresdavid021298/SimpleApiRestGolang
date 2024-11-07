package model

import "github.com/google/uuid"

type Motorcycle struct {
	Vehicle Vehicle
	Cc      int
	HasABS  bool
}

func NewMotorcycle(brand string, model string, amount string, color string, cc int, hasABS bool) Motorcycle {
	uuidGenerated := uuid.New()
	basedVehicle := Vehicle{Id: uuidGenerated, Brand: brand, Model: model, Amount: amount, Color: color}
	motorcycle := Motorcycle{HasABS: hasABS, Cc: cc, Vehicle: basedVehicle}
	return motorcycle
}
