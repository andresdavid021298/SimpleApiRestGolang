package util

import "ApiRest/model"

func CarCacheInit() []model.Car {
	cars := []model.Car{
		model.NewCar("Toyota", "Rav4 Híbrido", "178.000.000", "Silver", true, 29300),
		model.NewCar("Kia", "Cerato Forte", "30.000.000", "Silver", true, 167000),
		model.NewCar("Renault", "Sandero", "55.000.000", "Red", false, 0),
		model.NewCar("Toyota", "Corolla Cross", "129.900.000", "Gray", true, 41810),
	}
	return cars
}

func MotorcycleCacheInit() []model.Motorcycle {
	motorcycles := []model.Motorcycle{
		model.NewMotorcycle("KTM", "Adventure 250", "21.000.000", "Orange", 250, true),
		model.NewMotorcycle("Royal Enfield", "Hunter 350", "19.500.000", "Red", 350, false),
		model.NewMotorcycle("Suzuki", "V-Strom 650XT", "42.700.000", "White", 650, false),
	}
	return motorcycles
}
